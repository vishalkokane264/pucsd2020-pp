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
	"github.com/pucsd2020-pp/rest-api/repository/filemapping"
)

type FileMapping struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewFileMappingHandler(conn *sql.DB) *FileMapping {
	return &FileMapping{
		repo: filemapping.NewFileMappingRepository(conn),
	}
}

func (filemapping *FileMapping) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "filemapping/{file_id}", Func: filemapping.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "filemapping", Func: filemapping.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "filemapping/{file_id}", Func: filemapping.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "filemapping/{file_id}", Func: filemapping.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "filemapping", Func: filemapping.GetAll},
	}
}

func (filemapping *FileMapping) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "file_id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = filemapping.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (filemapping *FileMapping) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.FileMapping
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = filemapping.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (filemapping *FileMapping) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "file_id"), 10, 64)
	usr := model.FileMapping{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in user id for tracking update
		// usr.UpdatedBy = 0

		iUsr, err = filemapping.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.FileMapping)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (filemapping *FileMapping) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "file_id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = filemapping.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "File mapping deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (filemapping *FileMapping) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := filemapping.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
