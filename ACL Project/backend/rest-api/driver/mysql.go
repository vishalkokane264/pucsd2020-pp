package driver

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/pucsd2020-pp/rest-api/config"
	"github.com/pucsd2020-pp/rest-api/model"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_DRIVER_NAME   = "mysql"
	CONN_MAX_LIFETIME   = 30 * 60 * 60 // 30 day
	COLUMN_INGNORE_FLAG = "1"
	COLUMN_PRIMARY      = "primary"
)

func NewMysqlConnection(cfg config.MysqlConnection) (*sql.DB, error) {
	db, err := sql.Open(MYSQL_DRIVER_NAME, cfg.ConnString())
	if err != nil {
		log.Fatalf("Failed to open mysql connection: %v", err)
		return nil, err
	}

	if cfg.IdleConnection > 0 {
		db.SetMaxIdleConns(cfg.IdleConnection)
	}
	if cfg.MaxConnection > 0 {
		db.SetMaxOpenConns(cfg.MaxConnection)
	}
	db.SetConnMaxLifetime(time.Second * CONN_MAX_LIFETIME)

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping mysql: %v", err)
	}

	return db, err
}

// return the placeholder string with given count
func GetPlaceHolder(count int) string {
	if count > 0 {
		str := strings.Repeat("?, ", count)
		return str[:len(str)-2]
	}

	return ""
}

/**
 * Insert new row
 */
func Create(conn *sql.DB, object model.IModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	count := 0
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		// if value.IsNil() || COLUMN_INGNORE_FLAG == field.Tag.Get("autoincr") ||
		// 	COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
		// 	continue
		// }

		column := field.Tag.Get("column")
		columns = append(columns, column)
		params = append(params, value.Interface())
		count++
	}
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("INSERT INTO ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString("(")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(") VALUES(")
	queryBuffer.WriteString(GetPlaceHolder(count))
	queryBuffer.WriteString(");")

	query := queryBuffer.String()
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))

	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))

	}

	return result, errors.New(fmt.Sprintf("Success"))
}

/**
 * Update existing row with key column
 */
func UpdateById(conn *sql.DB, object model.IModel) error {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	keyColumns := []string{}
	var keyParams []interface{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		// if value.IsNil() ||
		// 	COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
		// 	continue
		// }

		column := field.Tag.Get("column")
		if COLUMN_PRIMARY == field.Tag.Get("key") {
			keyColumns = append(keyColumns, column+" = ?")
			keyParams = append(keyParams, value.Interface())

		} else {
			columns = append(columns, column+" = ?")
			params = append(params, value.Interface())
		}
	}

	for _, param := range keyParams {
		params = append(params, param)
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" SET ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" WHERE ")
	queryBuffer.WriteString(strings.Join(keyColumns, ", "))
	queryBuffer.WriteString(";")

	query := queryBuffer.String()
	//	log.Println("Update statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return errors.New(fmt.Sprintf("Failure"))
	}

	defer stmt.Close()
	_, err = stmt.Exec(params...)
	if nil != err {
		log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}

func GetUser(conn *sql.DB, object model.IModel, uname string) (model.IModel, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE username = ?")

	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, uname)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, errors.New(fmt.Sprintf("Failure"))
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Entry not found", uname))
	}
	fmt.Println("Printing object", object)

	return object, errors.New(fmt.Sprintf("Success"))
}
func DeleteUser(conn *sql.DB, object model.IModel, uname string) (int64, error) {
	var queryBuffer bytes.Buffer
	fmt.Println("Im deleting user", uname)

	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE username = ? ")
	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n", err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Failure"))
	}
	defer stmt.Close()
	result, err := stmt.Exec(uname)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("User not deleted successfully"))

	}
	fmt.Println("Affectderows")
	id, err := result.RowsAffected()
	if id != 0 {
		return id, errors.New(fmt.Sprintf("User deleted successfully"))
	}
	return -1, errors.New(fmt.Sprintf("User not deleted successfully"))
}

func UpdateUser(conn *sql.DB, object model.IModel, fname string, lname string, password string, username string) error {
	fmt.Println("Updating User")
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	keyColumns := []string{}
	var keyParams []interface{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		// if value.IsNil() ||
		// 	COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
		// 	continue
		// }

		column := field.Tag.Get("column")
		if COLUMN_PRIMARY == field.Tag.Get("key") {
			keyColumns = append(keyColumns, column+" = ?")
			keyParams = append(keyParams, value.Interface())

		} else {
			columns = append(columns, column+" = ?")
			params = append(params, value.Interface())
		}
	}

	for _, param := range keyParams {
		params = append(params, param)
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" SET ")
	queryBuffer.WriteString(" user_fname = ? , user_lname = ? , password = ?")
	queryBuffer.WriteString(" WHERE ")
	queryBuffer.WriteString("username = ?")
	queryBuffer.WriteString(";")

	query := queryBuffer.String()
	//	log.Println("Update statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return errors.New(fmt.Sprintf("Failure"))
	}

	defer stmt.Close()
	_, err = stmt.Exec(fname, lname, password, username)
	if nil != err {
		log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}
func GetById(conn *sql.DB, object model.IModel, id int64) (model.IModel, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE id = ?")

	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, id)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, errors.New(fmt.Sprintf("Failure"))
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Entry not found", id))
	}
	fmt.Println("Printing object", object)

	return object, errors.New(fmt.Sprintf("Success"))
}

func GetAll(conn *sql.DB, object model.IModel, limit, offset int64) ([]interface{}, error) {
	//	var obj []map[string]interface{}
	fmt.Println("Inside driver executing GetAll")
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	fmt.Println(object.String())
	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)

		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}
	fmt.Println(columns)
	var queryBuffer bytes.Buffer
	var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	// if 0 != limit && 0 != offset {
	// 	queryBuffer.WriteString(" LIMIT ? OFFSET ?")
	// 	params = append(params, limit)
	// 	params = append(params, offset)
	// }

	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, params...)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))
	}

	defer row.Close()
	objects := make([]interface{}, 0)
	records, err := row.Columns()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
		values := make([]interface{}, len(records))
		recordsWrite := make([]string, len(records))
		for index, _ := range records {
			values[index] = &recordsWrite[index]
		}

		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
		objects = append(objects, values)
	}
	return objects, errors.New(fmt.Sprintf("Success"))
}

func DeleteById(conn *sql.DB, object model.IModel, id int64) (sql.Result, error) {
	var queryBuffer bytes.Buffer
	fmt.Println("Insiede deletebyid")
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))
	}

	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return result, errors.New(fmt.Sprintf("Success"))
}

func SoftDeleteById(conn *sql.DB, object model.IModel, id int64) error {
	var queryBuffer bytes.Buffer
	fmt.Println("Inside softdelete")
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}

func ValidateLogin(conn *sql.DB, object model.IModel, uname string, pass string) (model.IModel, error) {
	fmt.Println("validating login")
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE username = ?")
	queryBuffer.WriteString(" AND password = ?")
	// fmt.Println("3.Object", object)

	query := queryBuffer.String()
	// log.Printf("ValidateLogin sql: %s\n", query)
	row, err := conn.Query(query, uname, pass)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, errors.New(fmt.Sprintf("Failure"))
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
	} else {
		return object, errors.New(fmt.Sprintf("User not exist"))
	}
	fmt.Println("Printing object", object)

	return object, errors.New(fmt.Sprintf("Success"))
}

func GetUserPermission(conn *sql.DB, object model.IModel, uname string) (model.IModel, error) {
	fmt.Println("validating userpermission")
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE username = ?")
	// fmt.Println("3.Object", object)

	query := queryBuffer.String()
	// log.Printf("ValidateLogin sql: %s\n", query)
	row, err := conn.Query(query, uname)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, errors.New(fmt.Sprintf("Failure"))
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
	} else {
		return object, errors.New(fmt.Sprintf("User not exist"))
	}
	fmt.Println("Returning success")
	return object, errors.New(fmt.Sprintf("Success"))
}

func GetGroupMembers(conn *sql.DB, object model.IModel, gr_name string) ([]interface{}, error) {
	//	var obj []map[string]interface{}
	fmt.Println("Inside driver executing GetGroupMembers")
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE user_grp_gr_name = ? or username = ? ")
	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, gr_name, gr_name)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	objects := make([]interface{}, 0)
	records, err := row.Columns()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
		values := make([]interface{}, len(records))
		recordsWrite := make([]string, len(records))
		for index, _ := range records {
			values[index] = &recordsWrite[index]
		}

		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
		objects = append(objects, values)
	}

	return objects, nil
}
func AddMemberToGroup(conn *sql.DB, object model.IModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	count := 0
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		// if value.IsNil() || COLUMN_INGNORE_FLAG == field.Tag.Get("autoincr") ||
		// 	COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
		// 	continue
		// }

		column := field.Tag.Get("column")
		columns = append(columns, column)
		params = append(params, value.Interface())
		count++
	}
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("INSERT INTO ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString("(")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(") VALUES(")
	queryBuffer.WriteString(GetPlaceHolder(count))
	queryBuffer.WriteString(");")

	query := queryBuffer.String()
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))

	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, errors.New(fmt.Sprintf("Failure"))

	}

	return result, errors.New(fmt.Sprintf("Success"))
}
func GetResources(conn *sql.DB, object model.IModel, ow_name string) ([]interface{}, error) {
	fmt.Println("Inside driver executing GetGroupMembers", ow_name)
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	fmt.Println(rValue, rType)
	columns := []string{}
	pointers := make([]interface{}, 0)
	fmt.Println(rValue, rType, columns, pointers)
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}
	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE name = ? ")
	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, ow_name)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	objects := make([]interface{}, 0)
	records, err := row.Columns()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
		values := make([]interface{}, len(records))
		recordsWrite := make([]string, len(records))
		for index, _ := range records {
			values[index] = &recordsWrite[index]
		}

		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
		objects = append(objects, values)
	}
	//	objects := make([]interface{}, 0)

	return objects, nil
}

func RevokeUserAccess(conn *sql.DB, object model.IModel, uname string, fid int64, oid string) (int64, error) {
	fmt.Println("uid:", uname, "fid:", fid, "oid:", oid, "Filename:", object.Table())
	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE id = ?")
	// queryBuffer.WriteString(" user_file_perm_file_id=? AND ")
	// queryBuffer.WriteString(" user_file_perm_given_by=? ")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Failure"))
	}

	defer stmt.Close()
	result, err := stmt.Exec(uname)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Data not deleted successfully"))

	}
	fmt.Println("Affectderows")
	id, err := result.RowsAffected()
	if id != 0 {
		return 1, errors.New(fmt.Sprintf("Data deleted successfully"))

	}
	return -1, errors.New(fmt.Sprintf("Data not deleted successfully"))
}

func RevokeGroupAccess(conn *sql.DB, object model.IModel, gr_name string, fid int64, oid string) (int64, error) {
	// fmt.Println("gid:", gid, "fid:", fid, "oid:", oid, "Filename:", object.Table())
	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE id = ? ")
	// queryBuffer.WriteString(" grp_file_perm_file_id=? AND ")
	// queryBuffer.WriteString(" grp_file_perm_given_by=? ")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Failure"))
	}

	defer stmt.Close()
	result, err := stmt.Exec(gr_name)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Data not deleted successfully"))

	}
	fmt.Println("Affectderows")
	id, err := result.RowsAffected()
	if id != 0 {
		return 1, errors.New(fmt.Sprintf("Data deleted successfully"))

	}
	return -1, errors.New(fmt.Sprintf("Data not deleted successfully"))
}

func ModifyUserAccess(conn *sql.DB, object model.IModel, username string, fid int64, given_by string, perm_type int64) error {
	fmt.Println("Modifying useraccess")
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	keyColumns := []string{}
	var keyParams []interface{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		// if value.IsNil() ||
		// 	COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
		// 	continue
		// }

		column := field.Tag.Get("column")
		if COLUMN_PRIMARY == field.Tag.Get("key") {
			keyColumns = append(keyColumns, column+" = ?")
			keyParams = append(keyParams, value.Interface())

		} else {
			columns = append(columns, column+" = ?")
			params = append(params, value.Interface())
		}
	}

	for _, param := range keyParams {
		params = append(params, param)
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" SET ")
	queryBuffer.WriteString(" user_file_perm_type=? ")
	queryBuffer.WriteString(" WHERE ")
	queryBuffer.WriteString("id=? AND user_file_perm_file_id=? AND user_file_perm_given_by=? ")
	queryBuffer.WriteString(";")

	query := queryBuffer.String()
	//	log.Println("Update statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return errors.New(fmt.Sprintf("Failure"))
	}

	defer stmt.Close()
	_, err = stmt.Exec(perm_type, username, fid, given_by)
	if nil != err {
		log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}

func ModifyGroupAccess(conn *sql.DB, object model.IModel, gr_name string, fid int64, given_by string, perm_type int64) error {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	keyColumns := []string{}
	var keyParams []interface{}

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		// if value.IsNil() ||
		// 	COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
		// 	continue
		// }

		column := field.Tag.Get("column")
		if COLUMN_PRIMARY == field.Tag.Get("key") {
			keyColumns = append(keyColumns, column+" = ?")
			keyParams = append(keyParams, value.Interface())

		} else {
			columns = append(columns, column+" = ?")
			params = append(params, value.Interface())
		}
	}

	for _, param := range keyParams {
		params = append(params, param)
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" SET ")
	queryBuffer.WriteString(" grp_file_perm_type=? ")
	queryBuffer.WriteString(" WHERE ")
	queryBuffer.WriteString("gr_name=? AND grp_file_perm_file_id=? AND grp_file_perm_given_by=? ")
	queryBuffer.WriteString(";")

	query := queryBuffer.String()
	//	log.Println("Update statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return errors.New(fmt.Sprintf("Failure"))
	}

	defer stmt.Close()
	_, err = stmt.Exec(perm_type, gr_name, fid, given_by)
	if nil != err {
		log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}

func DeleteGroupMember(conn *sql.DB, object model.IModel, uname string, gr_name string) (int64, error) {
	var queryBuffer bytes.Buffer
	fmt.Println(uname, gr_name)

	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE username = ? AND ")
	queryBuffer.WriteString(" user_grp_gr_name=?")
	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n", err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Failure"))
	}
	defer stmt.Close()
	result, err := stmt.Exec(uname, gr_name)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Data not deleted successfully"))
	}
	fmt.Println("Affectderows")
	id, err := result.RowsAffected()
	if id != 0 {
		return 1, errors.New(fmt.Sprintf("Data deleted successfully"))
	}
	return -1, errors.New(fmt.Sprintf("Data not deleted successfully"))
}
func DeleteGroup(conn *sql.DB, object model.IModel, gr_name string) (int64, error) {
	var queryBuffer bytes.Buffer
	fmt.Println("Printing group nameto delete", gr_name)
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE gr_name = ? ")
	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n", err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Failure"))
	}
	defer stmt.Close()
	result, err := stmt.Exec(gr_name)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return -1, errors.New(fmt.Sprintf("Data not deleted successfully"))
	}
	fmt.Println("Affectderows")
	id, err := result.RowsAffected()
	if id != 0 {
		return 1, errors.New(fmt.Sprintf("Data deleted successfully"))
	}
	return -1, errors.New(fmt.Sprintf("Data not deleted successfully"))
}

func DeleteUserPermission(conn *sql.DB, object model.IModel, uid int64, pid int64) (bool, error) {
	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE username = ? ")
	// queryBuffer.WriteString(" user_default_perm_user_type_id =?")
	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n", err.Error(), object.String(), query)
		return false, errors.New(fmt.Sprintf("Failure"))
	}

	defer stmt.Close()
	result, err := stmt.Exec(uid)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return false, errors.New(fmt.Sprintf("Data not deleted successfully"))

	}
	fmt.Println(result)
	if result != nil {
		return true, errors.New(fmt.Sprintf("Data deleted successfully"))

	}
	return false, errors.New(fmt.Sprintf("Data not deleted successfully"))
}

func GetFiles(conn *sql.DB, object model.IModel, gr_name int64) ([]interface{}, error) {
	//	var obj []map[string]interface{}
	fmt.Println("Inside driver executing GetGroupMembers")
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE parent = ? ")
	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, gr_name)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	objects := make([]interface{}, 0)
	records, err := row.Columns()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
		values := make([]interface{}, len(records))
		recordsWrite := make([]string, len(records))
		for index, _ := range records {
			values[index] = &recordsWrite[index]
		}

		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, errors.New(fmt.Sprintf("Failure"))
		}
		objects = append(objects, values)
	}

	return objects, nil
}

func GetFileData(conn *sql.DB, path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return "", errors.New(fmt.Sprintf("File not Open successfully"))
	}
	return string(data), errors.New(fmt.Sprintf("File Open successfully"))
}
