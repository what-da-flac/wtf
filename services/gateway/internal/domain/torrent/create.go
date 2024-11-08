package torrent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"
)

type Create struct {
	config     *environment.Config
	identifier interfaces2.Identifier
	timer      interfaces2.Timer
	repository interfaces2.Repository
	publisher  ifaces.Publisher
}

func NewCreate(
	config *environment.Config,
	identifier interfaces2.Identifier, repository interfaces2.Repository,
	timer interfaces2.Timer, sender ifaces.Publisher) *Create {
	return &Create{
		config:     config,
		identifier: identifier,
		repository: repository,
		publisher:  sender,
		timer:      timer,
	}
}

func (x *Create) validate(req *models.PostV1TorrentsMagnetsJSONRequestBody) error {
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

func (x *Create) Create(ctx context.Context, user *models.User, req *models.PostV1TorrentsMagnetsJSONRequestBody) error {
	if err := x.validate(req); err != nil {
		return err
	}
	now := x.timer.Now()
	for _, v := range *req.Urls {
		payload := &models.Torrent{
			Created:    now,
			Id:         x.identifier.UUIDv4(),
			MagnetLink: v,
			Status:     models.Pending,
			User:       user,
		}
		if err := x.repository.InsertTorrent(ctx, payload); err != nil {
			return err
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
