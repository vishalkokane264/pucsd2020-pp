package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *userRepository {
	return &userRepository{conn: conn}
}

func (user *userRepository) ValidateLogin(cntx context.Context, uname string, pass string) (interface{}, error) {
	obj := new(model.User)
	fmt.Println("Inside repository validateLogin")
	return driver.ValidateLogin(user.conn, obj, uname, pass)
}

func (user *userRepository) GetUser(cntx context.Context, username string) (interface{}, error) {
	obj := new(model.User)
	return driver.GetUser(user.conn, obj, username)
}
func (user *userRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.User)
	return driver.GetById(user.conn, obj, id)
}

func (user *userRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.User)
	result, err := driver.Create(user.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user *userRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.User)
	err := driver.UpdateById(user.conn, &usr)
	return obj, err
}

func (user *userRepository) DeleteUser(cntx context.Context, uname string) (int64, error) {
	obj := new(model.User)
	fmt.Println("Repo deleteuser", uname)
	fmt.Println("Inside delete repository user")
	return driver.DeleteUser(user.conn, obj, uname)
}

func (user *userRepository) UpdateUser(cntx context.Context, obj interface{}, fname string, lname string, password string, uname string) (interface{}, error) {
	usr := obj.(model.User)
	err := driver.UpdateUser(user.conn, &usr, fname, lname, password, uname)
	return obj, err
}

func (user *userRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.User{Id: id}
	return driver.SoftDeleteById(user.conn, obj, id)
}

func (user *userRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.User{}
	fmt.Println("Inside Repo user.go")
	return driver.GetAll(user.conn, obj, 0, 0)
}
