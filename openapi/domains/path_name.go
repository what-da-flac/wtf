package domains

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func NewPathName(name string) golang.PathName {
	v := golang.PathName(name)
	switch v {
	case golang.PathNameStore,
		golang.PathNameTemp:
	default:
		return golang.PathNameInvalid
	}
	return v
}
