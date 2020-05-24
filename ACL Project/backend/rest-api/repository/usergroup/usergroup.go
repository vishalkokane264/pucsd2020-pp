package usergroup

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type usergroupRepository struct {
	conn *sql.DB
}

func NewUserGroupRepository(conn *sql.DB) *usergroupRepository {
	return &usergroupRepository{conn: conn}
}
func (usergroup *usergroupRepository) AddMemberToGroup(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserGroup)
	result, err := driver.AddMemberToGroup(usergroup.conn, &usr)
	if nil != err {
		return 0, err
	}
	return result, nil
}

func (usergroup *usergroupRepository) GetGroupMembers(cntx context.Context, ug interface{}, gr_name string) ([]interface{}, error) {
	usr_grp := &model.UserGroup{}
	return driver.GetGroupMembers(usergroup.conn, usr_grp, gr_name)
}

func (usergroup *usergroupRepository) DeleteGroupMember(cntx context.Context, username string, gr_name string) (int64, error) {
	obj := new(model.UserGroup)
	fmt.Println("Inside repository ug")
	return driver.DeleteGroupMember(usergroup.conn, obj, username, gr_name)
}
