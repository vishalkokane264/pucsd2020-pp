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
	"github.com/pucsd2020-pp/rest-api/repository/userfilepermission"
)

type UserFilePermission struct {
	handler.HTTPHandler
	repo  repository.IRepository
	repo2 repository.AccessControl
}

func NewUserFilePermHandler(conn *sql.DB) *UserFilePermission {
	return &UserFilePermission{
		repo:  userfilepermission.NewUserPermissionRepository(conn),
		repo2: userfilepermission.NewUserPermissionRepository(conn),
	}
}

func (userfilepermission *UserFilePermission) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userfilepermission/{id}", Func: userfilepermission.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "userfilepermission", Func: userfilepermission.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "userfilepermission", Func: userfilepermission.ModifyUserAccess},
		// &handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "userfilepermission/{id}", Func: userfilepermission.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "userfilepermission/{username}", Func: userfilepermission.RevokeAccess},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userfilepermission", Func: userfilepermission.GetAll},
		///write path =Groups/{id}
	}
}

func (userfilepermission *UserFilePermission) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = userfilepermission.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userfilepermission *UserFilePermission) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.UserFilePermission
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = userfilepermission.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
func (userfilepermission *UserFilePermission) RevokeAccess(w http.ResponseWriter, r *http.Request) {
	var ufp model.UserFilePermission
	var ufp2 interface{}
	uname, _ := chi.URLParam(r, "username"), 10
	err := json.NewDecoder(r.Body).Decode(&ufp)
	fmt.Println("im printing user object", uname, ufp.File_id) //not sending other data except id as username

	fmt.Println("Inside Revokeaccesshandler uid:=", ufp.UserName, "fid:=", ufp.File_id, "oid:=", ufp.Given_by)
	ufp2, err = userfilepermission.repo2.RevokeAccess(r.Context(), uname, ufp.File_id, ufp.Given_by)
	handler.WriteJSONResponse(w, r, ufp2, http.StatusOK, err)

}

func (userfilepermission *UserFilePermission) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	fmt.Println(id)
	usr := model.UserFilePermission{}
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

		iUsr, err = userfilepermission.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.UserFilePermission)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

/*func (userfilepermission *UserFilePermission) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = userfilepermission.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "UserFilePermission deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}*/

func (userfilepermission *UserFilePermission) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := userfilepermission.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (userfilepermission *UserFilePermission) ModifyUserAccess(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	// id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.UserFilePermission{}
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

		iUsr, err = userfilepermission.repo2.ModifyUserAccess(r.Context(), usr, usr.UserName, usr.File_id, usr.Given_by, usr.Perm_type)
		if nil != err {
			break
		}
		usr = iUsr.(model.UserFilePermission)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
