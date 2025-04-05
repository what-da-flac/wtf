package torrent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"
)

type Create struct {
	config    *environment.Config
	timer     interfaces.Timer
	publisher ifaces.Publisher
}

func NewCreate(
	config *environment.Config,
	timer interfaces.Timer, sender ifaces.Publisher) *Create {
	return &Create{
		config:    config,
		publisher: sender,
		timer:     timer,
	}
}

func (x *Create) validate(req *golang.PostV1TorrentsMagnetsJSONRequestBody) error {
	if x.publisher == nil {
		return fmt.Errorf("missing publisher")
	}
	if req.Urls == nil {
		return fmt.Errorf("missing urls")
	}
	for _, u := range *req.Urls {
		if u == "" {
			return fmt.Errorf("missing url")
		}
	}
	return nil
}

func (x *Create) Create(ctx context.Context, user *golang.User, req *golang.PostV1TorrentsMagnetsJSONRequestBody) error {
	if err := x.validate(req); err != nil {
		return err
	}
	for _, v := range *req.Urls {
		now := x.timer.Now()
		payload := &golang.Torrent{
			Created:    now,
			MagnetLink: v,
			Status:     golang.Pending,
			User:       user,
		}
		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		if err := x.publisher.Publish(data); err != nil {
			return err
		}
	}
	return nil
}
