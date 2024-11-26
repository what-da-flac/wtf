package processors

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/what-da-flac/wtf/services/torrent-parser/internal/interfaces"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-parser/internal/parsing"
)

func Process(
	publisher ifaces.Publisher, downloader interfaces.S3Downloader,
	torrent *models.Torrent) error {
	// base dir must be /tmp since lambdas cannot write anywhere else
	// download torrent from s3
	file, err := os.CreateTemp(os.TempDir(), "_torrent")
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(file.Name()) }()
	if err = downloader.Download(file, env.BucketTorrentParsed.String(), torrent.Filename); err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}
	// extract metadata from torrent file
	metaInfo, err := torrentMetadata(file.Name())
	if err != nil {
		return err
	}
	// parse metadata into local torrent struct
	parsedTorrent, err := parsing.ParseTorrent(*metaInfo)
	if err != nil {
		return err
	}
	// save torrent file in s3
	parsed := parsedTorrent.ToDomain()
	torrent.Name = parsed.Name
	torrent.PieceCount = parsed.PieceCount
	torrent.PieceSize = parsed.PieceSize
	torrent.Privacy = parsed.Privacy
	torrent.TotalSize = parsed.TotalSize
	torrent.Files = parsed.Files
	data, err := json.Marshal(torrent)
	if err != nil {
		return err
	}
	return publisher.Publish(data)
}

func torrentMetadata(torrentFilename string) (*string, error) {
	cmd := exec.Command("transmission-show", torrentFilename)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return aws.String(string(output)), nil
}
