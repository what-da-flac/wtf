package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ihandlers"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/gateway/internal/domain/torrent"
)

func (x *Server) PostV1TorrentsMagnets(w http.ResponseWriter, r *http.Request) {
	payload := &models.PostV1TorrentsMagnetsJSONRequestBody{}
	if err := ihandlers.ReadJSON(r.Body, payload); err != nil {
		ihandlers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	ctx := x.context(r)
	user := x.ReadUserFromContext(ctx)
	if user == nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, fmt.Errorf("user not found"))
		return
	}
	if err := torrent.NewCreate(x.config, x.timer, x.publisher(env.QueueMagnetParser)).
		Create(ctx, user, payload); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusAccepted, nil, nil)
}

func (x *Server) GetV1Torrents(w http.ResponseWriter, r *http.Request, params models.GetV1TorrentsParams) {
	ctx := x.context(r)
	res, err := torrent.NewList(x.repository).List(ctx, params)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) GetV1TorrentsStatuses(w http.ResponseWriter, r *http.Request) {
	ctx := x.context(r)
	res := x.repository.ListTorrentStatuses(ctx)
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) GetV1TorrentsId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := x.context(r)
	res, err := torrent.NewLoad(x.repository).Load(ctx, id)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) PostV1TorrentsIdDownload(w http.ResponseWriter, r *http.Request, id string) {
	queue := env.QueueTorrentDownload
	pubDownload := x.publishers[queue]
	if pubDownload == nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, fmt.Errorf("no publisher found for: %s", queue))
		return
	}
	queue = env.QueueTorrentInfo
	pubInfo := x.publishers[queue]
	if pubInfo == nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, fmt.Errorf("no publisher found for: %s", queue))
		return
	}
	ctx := x.context(r)
	t, err := x.repository.SelectTorrent(ctx, id)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, err)
		return
	}
	t.Status = models.Queued
	data, err := json.Marshal(t)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	if err = pubInfo.Publish(data); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	if err = pubDownload.Publish(data); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	x.logger.Infof("send t with id: %s to download process", id)
}

func (x *Server) PutV1TorrentsIdStatusStatus(w http.ResponseWriter, r *http.Request, id string, status string) {
	var ok bool
	ctx := x.context(r)
	oldTorrent, err := x.repository.SelectTorrent(ctx, id)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, err)
		return
	}
	statuses := x.repository.ListTorrentStatuses(ctx)
	for _, st := range statuses {
		if st != status {
			continue
		}
		ok = true
		switch models.TorrentStatus(st) {
		case models.Downloaded:
			ihandlers.WriteResponse(w, http.StatusBadRequest, nil, fmt.Errorf("oldTorrent %s is downloaded, cannot change status", id))
			return
		}
	}
	if !ok {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, fmt.Errorf("status %s is not valid", status))
		return
	}
	oldTorrent.Status = models.TorrentStatus(status)
	if err = x.repository.UpdateTorrent(ctx, oldTorrent); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
}
