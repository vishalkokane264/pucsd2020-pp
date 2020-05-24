package groupsapi

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type groupsapiRepository struct {
	conn *sql.DB
}

func NewGroupsApiRepository(conn *sql.DB) *groupsapiRepository {
	return &groupsapiRepository{conn: conn}
}

func (groupsapi *groupsapiRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.GroupsApi)
	return driver.GetById(groupsapi.conn, obj, id)
}

func (groupsapi *groupsapiRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.GroupsApi)
	result, err := driver.Create(groupsapi.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (groupsapi *groupsapiRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.GroupsApi)
	err := driver.UpdateById(groupsapi.conn, &usr)
	return obj, err
}

func (groupsapi *groupsapiRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.GroupsApi{Id: id}
	return driver.SoftDeleteById(groupsapi.conn, obj, id)
}

func (groupsapi *groupsapiRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.GroupsApi{}
	return driver.GetAll(groupsapi.conn, obj, 0, 0)
}
func (groupsapi *groupsapiRepository) DeleteGroup(cntx context.Context, gr_name string) (int64, error) {
	obj := new(model.GroupsApi)
	fmt.Println("My group is", gr_name)
	return driver.DeleteGroup(groupsapi.conn, obj, gr_name)
}
