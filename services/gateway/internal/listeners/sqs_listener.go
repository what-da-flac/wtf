package listeners

import (
	"context"
	"time"

	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSListener struct {
	fn                  interfaces2.MessageReceiverFn
	logger              interfaces2.Logger
	maxNumberOfMessages int
	name                string
	svc                 *sqs.SQS
	uri                 string
	visibilityTimeout   time.Duration
	waitTime            time.Duration
}

func (x *SQSListener) Poll(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := x.poll(x.maxNumberOfMessages, x.fn); err != nil {
				return err
			}
		}
	}
}

func (x *SQSListener) poll(maxNumberOfMessages int, fn interfaces2.MessageReceiverFn) error {
	output, err := x.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(int64(maxNumberOfMessages)),
		QueueUrl:            aws.String(x.uri),
		VisibilityTimeout:   aws.Int64(int64(x.visibilityTimeout.Seconds())),
		WaitTimeSeconds:     aws.Int64(int64(x.waitTime.Seconds())),
	})
	if err != nil {
		return err
	}
	if len(output.Messages) == 0 {
		return nil
	}
	for _, msg := range output.Messages {
		if err = fn(*msg.Body); err != nil {
			x.logger.Errorf("error processing message with id: %s: %v", *msg.MessageId, err)
			return nil
		}
		if err = x.deleteMessage(msg.ReceiptHandle); err != nil {
			x.logger.Errorf("error deleting message with id: %s: %v", *msg.MessageId, err)
			return nil
		}
	}
	return nil
}

func (x *SQSListener) deleteMessage(handle *string) error {
	_, err := x.svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(x.uri),
		ReceiptHandle: handle,
	})
	return err
}

func (x *SQSListener) Name() string {
	return x.name
}

func NewSQSListener(sess *session.Session,
	fn interfaces2.MessageReceiverFn,
	name, uri string,
	logger interfaces2.Logger,
	visibilityTimeout, waitTime time.Duration,
	maxNumberOfMessages int,
) *SQSListener {
	svc := sqs.New(sess)
	return &SQSListener{
		fn:                  fn,
		logger:              logger,
		maxNumberOfMessages: maxNumberOfMessages,
		name:                name,
		svc:                 svc,
		uri:                 uri,
		visibilityTimeout:   visibilityTimeout,
		waitTime:            waitTime,
	}
}
