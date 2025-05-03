package domains

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func NewContentType(contentType string) golang.MediaInfoInputContentType {
	v := golang.MediaInfoInputContentType(contentType)
	switch v {
	case golang.MediaInfoInputContentTypeAudioFlac,
		golang.MediaInfoInputContentTypeAudioMp3,
		golang.MediaInfoInputContentTypeAudioMpeg,
		golang.MediaInfoInputContentTypeAudioXM4A:
	default:
		return golang.MediaInfoInputContentTypeInvalid
	}
	return v
}
