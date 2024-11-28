package interfaces

import (
	"github.com/what-da-flac/wtf/go-common/ifaces"
)

//go:generate moq -out ../../mocks/publisher.go -pkg mocks . Publisher
type Publisher interface {
	ifaces.Publisher
}
