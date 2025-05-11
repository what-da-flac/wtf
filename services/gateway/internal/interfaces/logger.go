package interfaces

import "github.com/what-da-flac/wtf/common/ifaces"

//go:generate moq -out ../../mocks/logger.go -pkg mocks . Logger
type Logger interface {
	ifaces.Logger
}
