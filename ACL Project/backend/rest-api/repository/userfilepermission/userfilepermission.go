package userfilepermission

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type userFilePermRepository struct {
	conn *sql.DB
}

func NewUserPermissionRepository(conn *sql.DB) *userFilePermRepository {
	return &userFilePermRepository{conn: conn}
}

func (userfilepermission *userFilePermRepository) RevokeAccess(cntx context.Context, uname string, fid int64, oid string) (int64, error) {
	obj := new(model.UserFilePermission)
	fmt.Println("Inside repository Revokeaccess")
	return driver.RevokeUserAccess(userfilepermission.conn, obj, uname, fid, oid)
}

func (userfilepermission *userFilePermRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.UserFilePermission)
	return driver.GetById(userfilepermission.conn, obj, id)
}

func (userfilepermission *userFilePermRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserFilePermission)
	result, err := driver.Create(userfilepermission.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	id = 1
	return id, nil
}

func (userfilepermission *userFilePermRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserFilePermission)
	err := driver.UpdateById(userfilepermission.conn, &usr)
	return obj, err
}

func (userfilepermission *userFilePermRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.UserFilePermission{}
	return driver.SoftDeleteById(userfilepermission.conn, obj, 1)
}

func (userfilepermission *userFilePermRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.UserFilePermission{}
	return driver.GetAll(userfilepermission.conn, obj, 0, 0)
}

func (userfilepermission *userFilePermRepository) ModifyUserAccess(cntx context.Context, obj interface{}, username string, fid int64, given_by string, perm_type int64) (interface{}, error) {
	usr := obj.(model.UserFilePermission)
	err := driver.ModifyUserAccess(userfilepermission.conn, &usr, username, fid, given_by, perm_type)
	return obj, err
}
