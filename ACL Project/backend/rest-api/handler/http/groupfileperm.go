package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/groupfilepermission"
)

type GroupFilePermission struct {
	handler.HTTPHandler
	repo  repository.IRepository
	repo2 repository.IGroupAccessControl
}

func NewGroupFilePermHandler(conn *sql.DB) *GroupFilePermission {
	return &GroupFilePermission{
		repo:  groupfilepermission.NewGroupFilePermissionRepository(conn),
		repo2: groupfilepermission.NewGroupFilePermissionRepository(conn),
	}
}

func (groupfilepermission *GroupFilePermission) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groupfilepermission/{id}", Func: groupfilepermission.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "groupfilepermission", Func: groupfilepermission.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "groupfilepermission", Func: groupfilepermission.ModifyGroupAccess},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "groupfilepermission/{id}", Func: groupfilepermission.RevokeGroupAccess},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groupfilepermission", Func: groupfilepermission.GetAll},
		///write path =Groups/{id}
	}
}

func (groupfilepermission *GroupFilePermission) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = groupfilepermission.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (groupfilepermission *GroupFilePermission) RevokeGroupAccess(w http.ResponseWriter, r *http.Request) {
	var gfp model.GroupFilePermission
	var gfp2 interface{}
	id, _ := chi.URLParam(r, "id"), 10

	err := json.NewDecoder(r.Body).Decode(&gfp)
	// fmt.Println("Inside RevokeGrpAccHand gid:=", gfp.Id, "fid:=", gfp.File_id, "oid:=", gfp.Given_by)
	gfp2, err = groupfilepermission.repo2.RevokeGroupAccess(r.Context(), id, gfp.File_id, gfp.Given_by)

	handler.WriteJSONResponse(w, r, gfp2, http.StatusOK, err)

}

func (groupfilepermission *GroupFilePermission) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.GroupFilePermission
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = groupfilepermission.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
func (groupfilepermission *GroupFilePermission) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	// id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.GroupFilePermission{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		// usr.Id = id
		if nil != err {
			break
		}

		// set logged in userrole id for tracking update
		//usr.UpdatedBy = 0

		iUsr, err = groupfilepermission.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.GroupFilePermission)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (groupfilepermission *GroupFilePermission) ModifyGroupAccess(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	// id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.GroupFilePermission{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		// usr.Id = id
		// if nil != err {
		// 	break
		// }

		// set logged in userrole id for tracking update
		//usr.UpdatedBy = 0

		iUsr, err = groupfilepermission.repo2.ModifyGroupAccess(r.Context(), usr, usr.Gr_Name, usr.File_id, usr.Given_by, usr.Perm_type)
		if nil != err {
			break
		}
		usr = iUsr.(model.GroupFilePermission)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (groupfilepermission *GroupFilePermission) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = groupfilepermission.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "groupfilepermission deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (groupfilepermission *GroupFilePermission) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := groupfilepermission.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
