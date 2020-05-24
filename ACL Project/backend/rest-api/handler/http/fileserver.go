package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/fileserver"
)

type FileServer struct {
	handler.HTTPHandler
	repo  repository.IRepository
	repo2 repository.IFileServer
}

func NewFileServerHandler(conn *sql.DB) *FileServer {
	return &FileServer{
		repo:  fileserver.NewFileServerRepository(conn),
		repo2: fileserver.NewFileServerRepository(conn),
	}
}

func (fileserver *FileServer) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "getpath/{id}", Func: fileserver.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "fileserver/{ownername}", Func: fileserver.GetResources},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "fileserver", Func: fileserver.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "getfiledata", Func: fileserver.GetFileData},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "fileserver/{file_id}", Func: fileserver.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "fileserver/{file_id}", Func: fileserver.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "fileserver", Func: fileserver.GetAll},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "getfiles/{level}", Func: fileserver.GetFiles},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "writedatatofile", Func: fileserver.WriteDataToFile},
	}
}

func (fileserver *FileServer) WriteDataToFile(w http.ResponseWriter, r *http.Request) {
	// var filedata string
	file_server := model.FileServer{}
	err := json.NewDecoder(r.Body).Decode(&file_server)

	fmt.Println(file_server.File_path, file_server.Owner_name)
	handler.WriteJSONResponse(w, r, writeToFile(file_server.File_path, file_server.Owner_name), http.StatusOK, err)
}
func writeToFile(path string, data string) string {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return "Cant open file"
	}
	defer file.Close()
	if _, err := file.WriteString(data); err != nil {
		return "Error while appending data"
	}
	return "Data saved"

}

func (fileserver *FileServer) GetFileData(w http.ResponseWriter, r *http.Request) {
	var filedata string
	file_server := model.FileServer{}
	err := json.NewDecoder(r.Body).Decode(&file_server)

	fmt.Println(file_server.File_path)
	filedata, err = fileserver.repo2.GetFileData(r.Context(), file_server.File_path)
	handler.WriteJSONResponse(w, r, filedata, http.StatusOK, err)
}

func (fileserver *FileServer) GetResources(w http.ResponseWriter, r *http.Request) {
	var usr2 interface{}
	file_server := model.FileServer{}
	err := json.NewDecoder(r.Body).Decode(&file_server)
	owner_name, _ := chi.URLParam(r, "ownername"), 10
	fmt.Println("Print uname", owner_name, err, usr2)
	usr2, err = fileserver.repo2.GetResources(r.Context(), file_server, owner_name)
	handler.WriteJSONResponse(w, r, usr2, http.StatusOK, err)
}

func (fileserver *FileServer) GetByID(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	iUsr, err = fileserver.repo.GetByID(r.Context(), id)

	handler.WriteJSONResponse(w, r, iUsr, http.StatusOK, err)
}

func (fileserver *FileServer) GetFiles(w http.ResponseWriter, r *http.Request) {
	var usr2 interface{}
	user_grp := model.FileServer{}
	err := json.NewDecoder(r.Body).Decode(&user_grp)
	gr_name, _ := strconv.ParseInt(chi.URLParam(r, "level"), 10, 64)
	fmt.Println("Print uname", gr_name)
	usr2, err = fileserver.repo2.GetFiles(r.Context(), user_grp, gr_name)
	handler.WriteJSONResponse(w, r, usr2, http.StatusOK, err)
}

func (fileserver *FileServer) Create(w http.ResponseWriter, r *http.Request) {
	var file model.FileServer

	// var oldfiledata interface{}
	err := json.NewDecoder(r.Body).Decode(&file)
	for {
		if nil != err {
			break
		}
		var flag = checkPathExist(file.File_path)
		if !flag {
		} else {
			_, err = fileserver.repo.Create(r.Context(), file)
			break
		}
	}
	handler.WriteJSONResponse(w, r, file, http.StatusOK, err)
}

func checkPathExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if nil != err {
			fmt.Println("cannot create file.. invalid path")
			return false
		}
		defer file.Close()
		return true
	}
	return false
}

func (fileserver *FileServer) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "file_id"), 10, 64)
	usr := model.FileServer{}
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

		iUsr, err = fileserver.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.FileServer)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (fileserver *FileServer) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "file_id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = fileserver.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "File deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (fileserver *FileServer) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := fileserver.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
