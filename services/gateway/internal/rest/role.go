package rest

import (
	"net/http"

	"github.com/what-da-flac/wtf/go-common/ihandlers"
	"github.com/what-da-flac/wtf/openapi/models"
	role2 "github.com/what-da-flac/wtf/services/gateway/internal/domain/role"
)

func (x *Server) GetV1Roles(w http.ResponseWriter, r *http.Request) {
	ctx := x.context(r)
	res, err := role2.NewList(x.repository).List(ctx)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) PostV1Roles(w http.ResponseWriter, r *http.Request) {
	payload := &models.PostV1RolesJSONRequestBody{}
	if err := ihandlers.ReadJSON(r.Body, payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := x.context(r)
	res, err := role2.NewCreate(x.identifier, x.repository).Save(ctx, payload)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusCreated, res, nil)
}

func (x *Server) DeleteV1RolesId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := x.context(r)
	if err := role2.NewDelete(x.repository).Delete(ctx, &models.Role{Id: id}); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusNoContent, nil, nil)
}

func (x *Server) GetV1RolesId(w http.ResponseWriter, r *http.Request, id string) {
	ctx := x.context(r)
	res, err := role2.NewLoad(x.repository).Load(ctx, id)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusNotFound, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) PutV1RolesId(w http.ResponseWriter, r *http.Request, id string) {
	payload := &models.RolePut{}
	if err := ihandlers.ReadJSON(r.Body, payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := x.context(r)
	if err := role2.NewUpdate(x.repository).Update(ctx, id, payload); err != nil {
		ihandlers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusNoContent, nil, nil)
}

func (x *Server) GetV1RolesRoleIdUsers(w http.ResponseWriter, r *http.Request, roleId string) {
	ctx := x.context(r)
	res, err := role2.NewListUsers(x.repository).List(ctx, roleId)
	if err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusOK, res, nil)
}

func (x *Server) DeleteV1RolesRoleIdUsersUserId(w http.ResponseWriter, r *http.Request, roleId string, userId string) {
	ctx := x.context(r)
	if err := role2.NewRemoveUser(x.repository).Remove(ctx,
		&models.Role{
			Id: roleId,
		}, &models.User{Id: userId}); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusNoContent, nil, nil)
}

func (x *Server) PutV1RolesRoleIdUsersUserId(w http.ResponseWriter, r *http.Request, roleId string, userId string) {
	ctx := x.context(r)
	if err := role2.NewAddUser(x.repository).Add(ctx,
		&models.Role{Id: roleId},
		&models.User{Id: userId}); err != nil {
		ihandlers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	ihandlers.WriteResponse(w, http.StatusNoContent, nil, nil)
}
