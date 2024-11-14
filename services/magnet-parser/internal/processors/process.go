package processors

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
)

func Process(publisher ifaces.Publisher, logger ifaces.Logger,
	sess *session.Session, config *env.Config, torrent *models.Torrent) error {
	// base dir must be /tmp since lambdas cannot write anywhere else
	baseDir := os.TempDir()
	// create torrent from magnet link
	torrentFile, err := createTorrentFile(baseDir, torrent.MagnetLink)
	if err != nil {
		logger.Error(err)
		return err
	}
	info, err := os.Stat(torrentFile)
	if err != nil {
		logger.Errorf("could not stat torrent file: %s", torrentFile)
		return err
	}
	logger.Infof("torrent file size: %v", info.Size())
	defer func() { _ = os.RemoveAll(torrentFile) }()
	// save torrent file in s3
	key := filepath.Base(torrentFile)
	// hash is always filename without extension,
	// the info may contain versions which makes this straightforward
	// and simpler
	torrent.Hash = strings.TrimSuffix(filepath.Base(torrentFile), filepath.Ext(torrentFile))
	// id is set to torrent hash, so we automatically avoid duplicated torrents in db
	torrent.Id = torrent.Hash
	torrent.Filename = key
	file, err := os.Open(torrentFile)
	if err != nil {
		logger.Errorf("could not open torrent file for reading: %s", torrentFile)
		return err
	}
	defer func() { _ = file.Close() }()
	if err = amazon.Upload(sess, file, env.BucketTorrentParsed.String(), key, amazon.Content{
		ContentDisposition: "attachment; filename=\"" + key + "\"",
		ContentLanguage:    "en",
		ContentLength:      info.Size(),
		ContentType:        "application/x-bittorrent",
	}); err != nil {
		logger.Errorf("could not upload torrent file to s3: %s", torrentFile)
		return err
	}
	// send resulting torrent struct to SQS
	data, err := json.Marshal(torrent)
	if err != nil {
		return err
	}
	return publisher.Publish(data)
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
