package groupfilepermission

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type groupFilePermRepository struct {
	conn *sql.DB
}

func NewGroupFilePermissionRepository(conn *sql.DB) *groupFilePermRepository {
	return &groupFilePermRepository{conn: conn}
}

func (groupfilepermission *groupFilePermRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.GroupFilePermission)
	return driver.GetById(groupfilepermission.conn, obj, id)
}

func (groupfilepermission *groupFilePermRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.GroupFilePermission)
	result, err := driver.Create(groupfilepermission.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	fmt.Println(id)
	// usr.Id = id
	return 1, nil
}

func (groupfilepermission *groupFilePermRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.GroupFilePermission)
	err := driver.UpdateById(groupfilepermission.conn, &usr)
	return obj, err
}

func (groupfilepermission *groupFilePermRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.GroupFilePermission{}
	return driver.SoftDeleteById(groupfilepermission.conn, obj, id)
}

func (groupfilepermission *groupFilePermRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.GroupFilePermission{}
	return driver.GetAll(groupfilepermission.conn, obj, 0, 0)
}

func (groupfilepermission *groupFilePermRepository) RevokeGroupAccess(cntx context.Context, gr_name string, fid int64, given_by string) (int64, error) {
	obj := new(model.GroupFilePermission)
	fmt.Println("Inside repository rga")
	return driver.RevokeGroupAccess(groupfilepermission.conn, obj, gr_name, fid, given_by)
}

func (groupfilepermission *groupFilePermRepository) ModifyGroupAccess(cntx context.Context, obj interface{}, gr_name string, fid int64, given_by string, perm_type int64) (interface{}, error) {
	usr := obj.(model.GroupFilePermission)
	err := driver.ModifyGroupAccess(groupfilepermission.conn, &usr, gr_name, fid, given_by, perm_type)
	return obj, err
}
