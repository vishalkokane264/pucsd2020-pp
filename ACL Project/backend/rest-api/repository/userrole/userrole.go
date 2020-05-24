package userrole

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type userRoleRepository struct {
	conn *sql.DB
}

func NewUserRoleRepository(conn *sql.DB) *userRoleRepository {
	return &userRoleRepository{conn: conn}
}

func (userrole *userRoleRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.UserRole)
	return driver.GetById(userrole.conn, obj, id)
}

func (userrole *userRoleRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserRole)
	result, err := driver.Create(userrole.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (userrole *userRoleRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.UserRole)
	err := driver.UpdateById(userrole.conn, &usr)
	return obj, err
}

func (userrole *userRoleRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.UserRole{Id: id}
	return driver.SoftDeleteById(userrole.conn, obj, id)
}

func (userrole *userRoleRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.UserRole{}
	return driver.GetAll(userrole.conn, obj, 0, 0)
}
