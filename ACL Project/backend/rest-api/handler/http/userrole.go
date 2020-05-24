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
	"github.com/pucsd2020-pp/rest-api/repository/userrole"
)

type UserRole struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewUserRoleHandler(conn *sql.DB) *UserRole {
	return &UserRole{
		repo: userrole.NewUserRoleRepository(conn),
	}
}

func (userrole *UserRole) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userrole/{id}", Func: userrole.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "userrole", Func: userrole.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "userrole/{id}", Func: userrole.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "userrole/{id}", Func: userrole.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userrole", Func: userrole.GetAll},
		///write path =Groups/{id}
	}
}

func (userrole *UserRole) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = userrole.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userrole *UserRole) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.UserRole
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = userrole.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userrole *UserRole) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.UserRole{}
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

		iUsr, err = userrole.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.UserRole)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userrole *UserRole) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = userrole.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "userrole deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (userrole *UserRole) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := userrole.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
