package rest

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/copier"

	"github.com/what-da-flac/wtf/gateway/internal/domain/role"
	"github.com/what-da-flac/wtf/gateway/internal/domain/user"
	"github.com/what-da-flac/wtf/go-common/ihandlers"
	"github.com/what-da-flac/wtf/openapi/models"
)

func (x *Server) PostV1UserList(w http.ResponseWriter, r *http.Request) {
	ctx := x.context(r)
	payload := &models.UserListParams{}
	if err := ihandlers.ReadJSON(r.Body, payload); err != nil {
		ihandlers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	res, err := user.NewList(x.repository).List(ctx, payload)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) PostV1UsersLogin(w http.ResponseWriter, r *http.Request) {
	ctx := x.context(r)
	payload := &models.PostV1UsersLoginJSONRequestBody{}
	if err := ihandlers.ReadJSON(r.Body, payload); err != nil {
		ihandlers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	res, err := user.NewLogin(x.identifier, x.repository, x.timer).Login(ctx, payload)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) PostV1Users(w http.ResponseWriter, r *http.Request) {
	ctx := x.context(r)
	payload := &models.UserPost{}
	if err := ihandlers.ReadJSON(r.Body, payload); err != nil {
		ihandlers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	res, err := user.NewCreate(x.identifier, x.repository, x.timer).Save(ctx, payload)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusCreated, res, nil)
}

func (x *Server) DeleteV1UsersId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := x.context(r)
	if err := user.NewDelete(x.repository).Delete(ctx, id); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusNoContent, nil, nil)
}

func (x *Server) GetV1UsersId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := x.context(r)
	res, err := user.NewLoad(x.repository).Load(ctx, &models.User{Id: id})
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) PutV1UsersId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := x.context(r)
	payload := &models.UserPut{}
	if err := ihandlers.ReadJSON(r.Body, payload); err != nil {
		ihandlers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	res, err := user.NewUpdate(x.repository, x.timer).Save(ctx, id, payload)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) GetV1UsersUserIdRoles(w http.ResponseWriter, r *http.Request, userId string) {
	ctx := x.context(r)
	res, err := role.NewListRoles(x.repository).List(ctx, &models.User{Id: userId})
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) GetV1UsersWhoami(w http.ResponseWriter, r *http.Request) {
	res := &models.UserLoginResponse{}
	ctx := r.Context()
	u := x.ReadUserFromContext(ctx)
	if u == nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, fmt.Errorf("user not found in context"))
	}
	if err := copier.Copy(res, u); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
	}
	res.Roles = ihandlers.RolesFromContext(ctx)
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}
