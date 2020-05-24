package fileserver

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type fileserverRepository struct {
	conn *sql.DB
}

func NewFileServerRepository(conn *sql.DB) *fileserverRepository {
	return &fileserverRepository{conn: conn}
}

func (fileserver *fileserverRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.FileServer)
	return driver.GetById(fileserver.conn, obj, id)
}

func (fileserver *fileserverRepository) GetFiles(cntx context.Context, ug interface{}, gr_name int64) ([]interface{}, error) {
	usr_grp := &model.FileServer{}
	return driver.GetFiles(fileserver.conn, usr_grp, gr_name)
}

func (fileserver *fileserverRepository) GetFileData(cntx context.Context, path string) (string, error) {
	fmt.Println("from repo", path)
	// return myStr, errors.New(fmt.Sprintf("No error"))
	return driver.GetFileData(fileserver.conn, path)
}

func (fileserver *fileserverRepository) GetResources(cntx context.Context, fs interface{}, owner_name string) ([]interface{}, error) {
	file_server := &model.FileServer{}
	fmt.Println("from repo", owner_name)
	return driver.GetResources(fileserver.conn, file_server, owner_name)
}

func (fileserver *fileserverRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.FileServer)
	result, err := driver.Create(fileserver.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (fileserver *fileserverRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.FileServer)
	err := driver.UpdateById(fileserver.conn, &usr)
	return obj, err
}

func (fileserver *fileserverRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.FileServer{Id: id}
	return driver.SoftDeleteById(fileserver.conn, obj, id)
}

func (fileserver *fileserverRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.FileServer{}
	return driver.GetAll(fileserver.conn, obj, 0, 0)
}
