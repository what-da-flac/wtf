package torrent

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type Create struct {
	config     *environment.Config
	identifier interfaces2.Identifier
	timer      interfaces2.Timer
	repository interfaces2.Repository
	sender     interfaces2.MessageSender
}

func NewCreate(
	config *environment.Config,
	identifier interfaces2.Identifier, repository interfaces2.Repository,
	timer interfaces2.Timer, sender interfaces2.MessageSender) *Create {
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
		// TODO: use rabbitmq instead
		//if err := x.sender.Send(x.config.SQS.TorrentMetadataUrl, payload); err != nil {
		//	return err
		//}
		return fmt.Errorf("not implemented")
	}
	return nil
}
