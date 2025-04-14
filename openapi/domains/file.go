package domains

import (
	"time"
)

type File struct {
	Id          string    `json:"id"`
	Filename    string    `json:"filename"`
	Created     time.Time `json:"created"`
	Length      int64     `json:"length"`
	ContentType string    `json:"content_type"`
	Status      string    `json:"status"`
}
