package rest

import (
	"fmt"
	"net/http"

	"github.com/what-da-flac/wtf/gateway/internal/domain/torrent"
	"github.com/what-da-flac/wtf/go-common/ihandlers"
	"github.com/what-da-flac/wtf/openapi/models"
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
	if err := torrent.NewCreate(x.config, x.identifier, x.repository, x.timer, x.messageSender).
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
