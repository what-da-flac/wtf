package torrent

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/gateway/internal/environment"

	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/openapi/models"
)

type Create struct {
	config     *environment.Config
	identifier interfaces.Identifier
	timer      interfaces.Timer
	repository interfaces.Repository
	sender     interfaces.MessageSender
}

func NewCreate(
	config *environment.Config,
	identifier interfaces.Identifier, repository interfaces.Repository,
	timer interfaces.Timer, sender interfaces.MessageSender) *Create {
	return &Create{
		config:     config,
		identifier: identifier,
		repository: repository,
		sender:     sender,
		timer:      timer,
	}
}

func (x *Create) validate(req *models.PostV1TorrentsMagnetsJSONRequestBody) error {
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
		if err := x.sender.Send(x.config.SQS.TorrentMetadataUrl, payload); err != nil {
			return err
		}
	}
	return nil
}
