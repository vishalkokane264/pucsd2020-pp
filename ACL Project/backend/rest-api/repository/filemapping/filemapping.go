package filemapping

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type filemappingRepository struct {
	conn *sql.DB
}

func NewFileMappingRepository(conn *sql.DB) *filemappingRepository {
	return &filemappingRepository{conn: conn}
}

func (filemapping *filemappingRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.FileMapping)
	return driver.GetById(filemapping.conn, obj, id)
}

func (filemapping *filemappingRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.FileMapping)
	result, err := driver.Create(filemapping.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (filemapping *filemappingRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.FileMapping)
	err := driver.UpdateById(filemapping.conn, &usr)
	return obj, err
}

func (filemapping *filemappingRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.FileMapping{Id: id}
	return driver.SoftDeleteById(filemapping.conn, obj, id)
}

func (filemapping *filemappingRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.FileMapping{}
	return driver.GetAll(filemapping.conn, obj, 0, 0)
}
