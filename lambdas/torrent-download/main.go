package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/lambdas/torrent-download/internal/environment"
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
	//baseDir := os.TempDir()
	// create torrent from magnet link
	return nil
}
