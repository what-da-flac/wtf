package interfaces

import "github.com/what-da-flac/wtf/go-common/ifaces"

//go:generate moq -out ../../mocks/logger.go -pkg mocks . Logger
type Logger interface {
	ifaces.Logger
}
