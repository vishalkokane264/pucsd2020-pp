package userpermission

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type userpermissionRepository struct {
	conn *sql.DB
}

func NewUserPermissionRepository(conn *sql.DB) *userpermissionRepository {
	return &userpermissionRepository{conn: conn}
}

func (userpermission *userpermissionRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.UserPermission)
	return driver.GetById(userpermission.conn, obj, id)
}

func (user *userpermissionRepository) GetUserPermission(cntx context.Context, uname string) (interface{}, error) {
	obj := new(model.UserPermission)
	return driver.GetUserPermission(user.conn, obj, uname)
}
func (userpermission *userpermissionRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserPermission)
	result, err := driver.Create(userpermission.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (userpermission *userpermissionRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserPermission)
	err := driver.UpdateById(userpermission.conn, &usr)
	return obj, err
}

func (userpermission *userpermissionRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.UserPermission{Id: id}
	return driver.SoftDeleteById(userpermission.conn, obj, id)
}

func (userpermission *userpermissionRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.UserPermission{}
	return driver.GetAll(userpermission.conn, obj, 0, 0)
}

func (userpermission *userpermissionRepository) DeleteUserPermission(cntx context.Context, uid int64, pid int64) (bool, error) {
	obj := new(model.UserPermission)
	return driver.DeleteUserPermission(userpermission.conn, obj, uid, pid)
}
