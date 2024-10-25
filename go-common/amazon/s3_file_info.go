package amazon

import (
	"encoding/json"
	"time"
)

type S3FileInfo struct {
	Bucket        string    `json:"bucket"`
	ContentLength int64     `json:"content_length"`
	ContentType   string    `json:"content_type"`
	Key           string    `json:"key"`
	LastModified  time.Time `json:"last_modified"`
}

func (x *S3FileInfo) String() string {
	data, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(data)
}
