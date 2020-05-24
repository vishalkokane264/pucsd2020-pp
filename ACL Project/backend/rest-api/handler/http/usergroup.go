package http

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/go-chi/chi"
	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/usergroup"
)

type UserGroup struct {
	handler.HTTPHandler
	repo2 repository.IUserGroup
}

func NewUserGroupHandler(conn *sql.DB) *UserGroup {
	return &UserGroup{
		repo2: usergroup.NewUserGroupRepository(conn),
	}
}

func (usergroup *UserGroup) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "usergroup/{gr_name}", Func: usergroup.GetGroupMembers},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "usergroup", Func: usergroup.AddMemberToGroup},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "usergroup", Func: usergroup.DeleteGroupMember},
	}
}

func (usergroup *UserGroup) AddMemberToGroup(w http.ResponseWriter, r *http.Request) {
	var usr model.UserGroup
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = usergroup.repo2.AddMemberToGroup(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (usergroup *UserGroup) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	var usr2 interface{}
	user_grp := model.UserGroup{}
	err := json.NewDecoder(r.Body).Decode(&user_grp)
	gr_name, _ := chi.URLParam(r, "gr_name"), 10
	fmt.Println("Print uname", gr_name)
	usr2, err = usergroup.repo2.GetGroupMembers(r.Context(), user_grp, gr_name)
	handler.WriteJSONResponse(w, r, usr2, http.StatusOK, err)
}

func (usergroup *UserGroup) DeleteGroupMember(w http.ResponseWriter, r *http.Request) {
	var ug model.UserGroup
	var ug2 interface{}
	err := json.NewDecoder(r.Body).Decode(&ug)
	fmt.Println(ug.User_Name, ug.Gr_Name)
	for {
		if nil != err {
			break
		}
		// fmt.Println("Inside RevokeGrpAccHand gid:=", gfp.Id, "fid:=", gfp.File_id, "oid:=", gfp.Given_by)
		ug2, err = usergroup.repo2.DeleteGroupMember(r.Context(), ug.User_Name, ug.Gr_Name)
		break
	}

	handler.WriteJSONResponse(w, r, ug2, http.StatusOK, err)

}
