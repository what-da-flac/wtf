package senders

import (
	"encoding/json"

	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type MessageSender struct {
	identifier interfaces2.Identifier
	logger     interfaces2.Logger
	svc        *sqs.SQS
}

func (x *MessageSender) Send(queueUri string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	bodyStr := string(body)
	if _, err = x.svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:            aws.String(bodyStr),
		MessageDeduplicationId: aws.String(x.identifier.UUIDv4()),
		MessageGroupId:         aws.String(x.identifier.UUIDv4()),
		QueueUrl:               aws.String(queueUri),
	}); err != nil {
		x.logger.Errorf("failed to send message with body: %v error: %s", bodyStr, err)
		return err
	}
	x.logger.Infof("successfully sent message with body: %s", bodyStr)
	return nil
}

func NewMessageSender(sess *session.Session, logger interfaces2.Logger,
	identifier interfaces2.Identifier) *MessageSender {
	svc := sqs.New(sess)
	return &MessageSender{
		identifier: identifier,
		logger:     logger,
		svc:        svc,
	}
}
