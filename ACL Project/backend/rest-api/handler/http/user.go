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
	"github.com/pucsd2020-pp/rest-api/repository/user"
)

type User struct {
	handler.HTTPHandler
	repo  repository.IRepository
	repo2 repository.ILogin
}

func NewUserHandler(conn *sql.DB) *User {
	return &User{
		repo:  user.NewUserRepository(conn),
		repo2: user.NewUserRepository(conn),
	}
}

func (user *User) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user/{username}", Func: user.GetUser},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "user", Func: user.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "user/{username}", Func: user.UpdateUser},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "user/{username}", Func: user.DeleteUser},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user", Func: user.GetAll},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "loginmaster", Func: user.ValidateLogin},
	}
}

func (user *User) ValidateLogin(w http.ResponseWriter, r *http.Request) {
	var usr model.User
	var usr2 interface{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		fmt.Println("Inside validate login: id:=", usr.Id, "pass:=", usr.Password)
		usr2, err = user.repo2.ValidateLogin(r.Context(), usr.UserName, usr.Password)
		break
	}

	handler.WriteJSONResponse(w, r, usr2, http.StatusOK, err)
}
func (user *User) GetUser(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := chi.URLParam(r, "username"), 10
	iUsr, err := user.repo2.GetUser(r.Context(), id)
	handler.WriteJSONResponse(w, r, iUsr, http.StatusOK, err)
}

func (user *User) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	fmt.Println("Entered id is", id)
	for {
		if nil != err {
			break
		}

		usr, err = user.repo.GetByID(r.Context(), id)
		fmt.Println("Response is", usr)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *User) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = user.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *User) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	usr := model.User{}
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

		iUsr, err = user.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.User)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *User) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	fmt.Println("Endsdd")
	id, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = user.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "User deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}
func (user *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var usr model.User
	var usr2 interface{}
	username, _ := chi.URLParam(r, "username"), 10
	fmt.Println("Endsdd", username)

	err := json.NewDecoder(r.Body).Decode(&usr)
	// fmt.Println("Inside RevokeGrpAccHand gid:=", gfp.Id, "fid:=", gfp.File_id, "oid:=", gfp.Given_by)
	usr2, err = user.repo2.DeleteUser(r.Context(), username)

	handler.WriteJSONResponse(w, r, usr2, http.StatusOK, err)

}

func (user *User) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside http user.go")
	usrs, err := user.repo.GetAll(r.Context())
	fmt.Println("returning from http user.go")
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
func (user *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	// id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.User{}
	uname, _ := chi.URLParam(r, "username"), 10

	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		iUsr, err = user.repo2.UpdateUser(r.Context(), usr, usr.FirstName, usr.LastName, usr.Password, uname)
		if nil != err {
			break
		}
		usr = iUsr.(model.User)
		usr.UserName = uname
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
