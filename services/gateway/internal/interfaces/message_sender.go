package interfaces

//go:generate moq -out ../../mocks/message_sender.go -pkg mocks . MessageSender
type MessageSender interface {
	Send(queueUrl string, payload any) error
}
