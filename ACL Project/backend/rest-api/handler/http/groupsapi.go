package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/groupsapi"
)

type GroupsApi struct {
	handler.HTTPHandler
	repo  repository.IRepository
	repo2 repository.IGroup
}

func NewGroupsApiHandler(conn *sql.DB) *GroupsApi {
	return &GroupsApi{
		repo:  groupsapi.NewGroupsApiRepository(conn),
		repo2: groupsapi.NewGroupsApiRepository(conn),
	}
}

func (groupsapi *GroupsApi) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "groups", Func: groupsapi.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "groups/{gr_name}", Func: groupsapi.DeleteGroup},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groups", Func: groupsapi.GetAll},
		// &handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groups/{id}", Func: groupsapi.GetByID},
		// &handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "groups/{id}", Func: groupsapi.Update},
		// &handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "groups/{id}", Func: groupsapi.Delete},
		///write path =Groups/{id}
	}
}

func (groupsapi *GroupsApi) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = groupsapi.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (groupsapi *GroupsApi) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.GroupsApi
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = groupsapi.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (groupsapi *GroupsApi) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.GroupsApi{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in userrole id for tracking update
		//usr.UpdatedBy = 0

		iUsr, err = groupsapi.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.GroupsApi)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (groupsapi *GroupsApi) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = groupsapi.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "groups deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (groupsapi *GroupsApi) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := groupsapi.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (groupsapi *GroupsApi) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	var gr2 interface{}
	gr_name, _ := chi.URLParam(r, "gr_name"), 10
	fmt.Println("My group name is", gr_name)

	// fmt.Println("Inside RevokeGrpAccHand gid:=", gfp.Id, "fid:=", gfp.File_id, "oid:=", gfp.Given_by)
	gr2, err := groupsapi.repo2.DeleteGroup(r.Context(), gr_name)

	handler.WriteJSONResponse(w, r, gr2, http.StatusOK, err)

}
