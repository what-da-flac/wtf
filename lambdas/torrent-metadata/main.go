package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/identifiers"
	"github.com/what-da-flac/wtf/lambdas/torrent-metadata/internal/environment"
	"github.com/what-da-flac/wtf/lambdas/torrent-metadata/internal/parsing"
	"github.com/what-da-flac/wtf/openapi/models"
)

func main() {
	lambda.Start(handler)
}

// handler converts magnet link to torrent file, extracts info and files,
// and sends sqs message with the information.
func handler(_ context.Context, sqsEvent *events.SQSEvent) error {
	config := environment.New()
	awsSession := amazon.NewAWSSessionFromEnvironment()
	if err := awsSession.Build(); err != nil {
		return err
	}
	sess := awsSession.Session()
	// loop over messages received
	for _, record := range sqsEvent.Records {
		payload := &models.Torrent{}
		if err := json.Unmarshal([]byte(record.Body), payload); err != nil {
			log.Println(err)
			return err
		}
		if err := process(sess, config, payload); err != nil {
			log.Println(err)
			// TODO: if above fails, send a message to another queue to deal with failed torrents
			// ignoring it for the time being
			return nil
		}
	}
	return nil
}

func process(sess *session.Session, config *environment.Config, torrent *models.Torrent) error {
	// base dir must be /tmp since lambdas cannot write anywhere else
	baseDir := os.TempDir()
	// create torrent from magnet link
	torrentFile, err := createTorrentFile(baseDir, torrent.MagnetLink)
	if err != nil {
		return err
	}
	info, err := os.Stat(torrentFile)
	if err != nil {
		log.Println("could not stat torrent file:", torrentFile)
		return err
	}
	log.Println("torrent file size:", info.Size())
	defer func() { _ = os.RemoveAll(torrentFile) }()
	// extract metadata from torrent file
	metaInfo, err := torrentMetadata(torrentFile)
	if err != nil {
		return err
	}
	log.Println("metaInfo:", *metaInfo)
	// parse metadata into local torrent struct
	parsedTorrent, err := parsing.ParseTorrent(*metaInfo)
	if err != nil {
		return err
	}
	// save torrent file in s3
	key := filepath.Base(torrentFile)
	payload := parsedTorrent.ToDomain()
	// hash is always filename without extension,
	// the info may contain versions which makes this straightforward
	// and simpler
	payload.Hash = strings.TrimSuffix(filepath.Base(torrentFile), filepath.Ext(torrentFile))
	payload.Id = torrent.Id
	payload.User = torrent.User
	payload.Filename = key
	payload.MagnetLink = torrent.MagnetLink
	payload.Status = models.Parsed
	file, err := os.Open(torrentFile)
	if err != nil {
		log.Println("could not open torrent file for reading:", torrentFile)
		return err
	}
	defer func() { _ = file.Close() }()
	if err = amazon.Upload(sess, file, config.S3TorrentParsedBucket, key, amazon.Content{
		ContentDisposition: "attachment; filename=\"" + key + "\"",
		ContentLanguage:    "en",
		ContentLength:      info.Size(),
		ContentType:        "application/x-bittorrent",
	}); err != nil {
		log.Println("could not upload torrent file to s3:", torrentFile)
		return err
	}
	// send resulting torrent struct to SQS
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	svc := sqs.New(sess)
	id := identifiers.NewIdentifier().UUIDv4()
	input := &sqs.SendMessageInput{
		MessageBody:            aws.String(string(data)),
		MessageDeduplicationId: aws.String(id),
		MessageGroupId:         aws.String(id),
		QueueUrl:               aws.String(config.SQSTorrentParsedURL),
	}
	_, err = svc.SendMessage(input)
	return err
}

func createTorrentFile(baseDir, magnet string) (string, error) {
	var res string
	cmd := exec.Command(
		"aria2c",
		"--bt-metadata-only=true",
		"--bt-save-metadata=true",
		"--dir="+baseDir,
		magnet,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	log.Println("output:", string(output))
	_ = filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		log.Println("found file:", d.Name())
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(d.Name(), ".torrent") {
			res = path
			return io.EOF
		}
		return nil
	})
	if res != "" {
		log.Println("identified torrent file at path:", res)
		return res, nil
	}
	return "", fmt.Errorf("unable to find torrent file")
}

func torrentMetadata(torrentFilename string) (*string, error) {
	cmd := exec.Command("transmission-show", torrentFilename)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return aws.String(string(output)), nil
}
