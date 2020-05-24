package repository

import (
	"context"
)

type IRepository interface {
	GetByID(context.Context, int64) (interface{}, error)
	Create(context.Context, interface{}) (interface{}, error)
	Update(context.Context, interface{}) (interface{}, error)
	Delete(context.Context, int64) error
	GetAll(context.Context) ([]interface{}, error)
}

type ILogin interface {
	ValidateLogin(context.Context, string, string) (interface{}, error)
	GetUser(context.Context, string) (interface{}, error)
	UpdateUser(context.Context, interface{}, string, string, string, string) (interface{}, error)
	DeleteUser(context.Context, string) (int64, error)
}
type IUserPermission interface {
	// ModifyUserPermission(context.Context, interface{}, int64, int64, int64) (interface{}, error)
	DeleteUserPermission(context.Context, int64, int64) (bool, error)
	GetUserPermission(context.Context, string) (interface{}, error)
}

type IUserGroup interface {
	AddMemberToGroup(context.Context, interface{}) (interface{}, error)
	GetGroupMembers(context.Context, interface{}, string) ([]interface{}, error)
	DeleteGroupMember(context.Context, string, string) (int64, error)
}

type IGroup interface {
	DeleteGroup(context.Context, string) (int64, error)
}

type IFileServer interface {
	GetResources(context.Context, interface{}, string) ([]interface{}, error)
	GetFiles(context.Context, interface{}, int64) ([]interface{}, error)
	GetFileData(context.Context, string) (string, error)
}

type AccessControl interface {
	RevokeAccess(context.Context, string, int64, string) (int64, error)
	ModifyUserAccess(context.Context, interface{}, string, int64, string, int64) (interface{}, error)
}

type IGroupAccessControl interface {
	RevokeGroupAccess(context.Context, string, int64, string) (int64, error)
	ModifyGroupAccess(context.Context, interface{}, string, int64, string, int64) (interface{}, error)
}

type Repository struct {
}

func (repo *Repository) GetFileData(cntx context.Context, path string) (data string, err error) {
	return
}

func (repo *Repository) DeleteGroup(cntx context.Context, gr_name string) (stat int64, err error) {
	return
}

func (repo *Repository) GetUser(cntx context.Context, username string) (obj interface{}, err error) {
	return
}

func (repo *Repository) DeleteUser(cntx context.Context, uname string) (rows int64, err error) {
	return
}
func (repo *Repository) UpdateUser(cntx context.Context, obj interface{}, fname string, lname string, passwd string, uname string) (uobj interface{}, err error) {
	return
}

func (repo *Repository) DeleteUserPermission(cntx context.Context, uid int64, pid int64) (boolean bool, err error) {
	return
}

func (repo *Repository) AddMemberToGroup(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}
func (repo *Repository) GetGroupMembers(cntx context.Context, obj interface{}, gr_name string) (obj1 []interface{}, err error) {
	return
}

func (repo *Repository) DeleteGroupMember(cntx context.Context, uname string, gr_name string) (stat int64, err error) {
	return
}

func (repo *Repository) RevokeAccess(cntx context.Context, uid int64, fid int64, oid string) (id int64, err error) {
	return
}
func (repo *Repository) ModifyUserAccess(cntx context.Context, obj interface{}, id int64, fid int64, oid string, ptype int64) (uobj interface{}, err error) {
	return
}

func (repo *Repository) RevokeGroupAccess(cntx context.Context, gid int64, fid int64, oid int64) (id int64, err error) {
	return
}
func (repo *Repository) ModifyGroupAccess(cntx context.Context, obj interface{}, id int64, fid int64, oid int64, ptype int64) (uobj interface{}, err error) {
	return
}

func (repo *Repository) GetUserPermission(cntx context.Context, uname string) (obj interface{}, err error) {
	return
}

func (repo *Repository) ValidateLogin(cntx context.Context, uname string, pass string) (obj interface{}, err error) {
	return
}

func (repo *Repository) GetResources(cntx context.Context, obj interface{}, gr_name string) (obj1 []interface{}, err error) {
	return
}

func (repo *Repository) GetByID(cntx context.Context, id int64) (obj interface{}, err error) {
	return
}
func (repo *Repository) GetFiles(cntx context.Context, obj interface{}, gr_name int64) (obj1 []interface{}, err error) {
	return
}

func (repo *Repository) Create(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (repo *Repository) Update(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

func (repo *Repository) Delete(cntx context.Context, id int64) (deleted bool, err error) {
	return
}

func (repo *Repository) GetAll(cntx context.Context) (obj []interface{}, err error) {
	return
}
