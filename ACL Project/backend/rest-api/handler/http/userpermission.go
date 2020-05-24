package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"fmt"

	"github.com/go-chi/chi"
	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/userpermission"
)

type UserPermission struct {
	handler.HTTPHandler
	repo  repository.IRepository
	repo2 repository.IUserPermission
}

func NewUserPermissionHandler(conn *sql.DB) *UserPermission {
	return &UserPermission{
		repo:  userpermission.NewUserPermissionRepository(conn),
		repo2: userpermission.NewUserPermissionRepository(conn),
	}
}

func (userpermission *UserPermission) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userpermission/{username}", Func: userpermission.GetUserPermission},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "userpermission", Func: userpermission.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "userpermission/{id}", Func: userpermission.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "userpermission/{username}", Func: userpermission.DeleteUserPermission},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userpermission", Func: userpermission.GetAll},
		///write path =Groups/{id}
	}
}

func (userpermission *UserPermission) GetUserPermission(w http.ResponseWriter, r *http.Request) {
	var userp model.UserPermission
	var iuserp interface{}
	uname, _ := chi.URLParam(r, "username"), 10
	err := json.NewDecoder(r.Body).Decode(&userp)
	fmt.Println("Given value", uname)
	fmt.Println("Calling Get Permission")
	iuserp, err = userpermission.repo2.GetUserPermission(r.Context(), uname)
	handler.WriteJSONResponse(w, r, iuserp, http.StatusOK, err)
}
func (userpermission *UserPermission) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.UserPermission
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = userpermission.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userpermission *UserPermission) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.UserPermission{}
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

		iUsr, err = userpermission.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.UserPermission)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userpermission *UserPermission) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = userpermission.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "User Permission deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (userpermission *UserPermission) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := userpermission.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (userpermission *UserPermission) DeleteUserPermission(w http.ResponseWriter, r *http.Request) {
	var up model.UserPermission
	var up2 interface{}
	uname, _ := strconv.ParseInt(chi.URLParam(r, "username"), 10, 64)
	err := json.NewDecoder(r.Body).Decode(&up)
	for {
		if nil != err {
			break
		}
		// fmt.Println("Inside Revokeaccesshandler uid:=", ufp.Id, "fid:=", ufp.File_id, "oid:=", ufp.Given_by)
		up2, err = userpermission.repo2.DeleteUserPermission(r.Context(), uname, up.Type_id)
		break
	}

	handler.WriteJSONResponse(w, r, up2, http.StatusOK, err)

}
