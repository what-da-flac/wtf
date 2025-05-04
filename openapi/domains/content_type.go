package domains

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func NewContentType(contentType string) golang.ContentType {
	v := golang.ContentType(contentType)
	switch v {
	case golang.ContentTypeAudioflac,
		golang.ContentTypeAudiomp3,
		golang.ContentTypeAudiompeg,
		golang.ContentTypeAudioxM4a:
	default:
		return golang.ContentTypeInvalid
	}
	return v
}
