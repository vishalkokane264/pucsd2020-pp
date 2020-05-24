package model

type User struct {
	Id        int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	FirstName string `json:"user_fname" column:"user_fname"`
	LastName  string `json:"user_lname" column:"user_lname"`
	UserName  string `json:"username" column:"username"`
	Password  string `json:"password" column:"password"`
}

func (user *User) Table() string {
	return "user"
}

func (user *User) String() string {
	return Stringify(user)
}

type UserRole struct {
	Id        int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	Role_name string `json:"role_name" column:"role_name"`
}

func (userrole *UserRole) Table() string {
	return "user_role"
}

func (userrole *UserRole) String() string {
	return Stringify(userrole)
}

type GroupsApi struct {
	Id          int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	Group_name  string `json:"group_name" column:"gr_name"`
	Group_owner string `json:"group_owner" column:"gr_owner"`
}

func (groupsapi *GroupsApi) Table() string {
	return "groups"
}

func (groupsapi *GroupsApi) String() string {
	return Stringify(groupsapi)
}

type UserPermission struct {
	Id       int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	UserName string `json:"username" column:"username"`
	Type_id  int64  `json:"type_id" column:"user_default_perm_user_type_id"`
}

func (userpermission *UserPermission) Table() string {
	return "user_default_perm"
}

func (userpermission *UserPermission) String() string {
	return Stringify(userpermission)
}

type FileServer struct {
	Id         int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	File_path  string `json:"file_path" column:"file_path"`
	Owner_name string `json:"owner_name" column:"owner_name"`
}

func (fileserver *FileServer) Table() string {
	return "file_server"
}

func (fileserver *FileServer) String() string {
	return Stringify(fileserver)
}

type FileMapping struct {
	Id        int64 `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	SubFileId int64 `json:"subfile_id" column:"subfile_id"`
}

func (filemapping *FileMapping) Table() string {
	return "file_mapping"
}

func (filemapping *FileMapping) String() string {
	return Stringify(filemapping)
}

type UserGroup struct {
	User_Name string `json:"user_name,omitempty" key:"primary"column:"username"`
	Gr_Name   string `json:"gr_name" column:"user_grp_gr_name"`
}

func (usergroup *UserGroup) Table() string {
	return "user_group"
}

func (usergroup *UserGroup) String() string {
	return Stringify(usergroup)
}

type UserFilePermission struct {
	Id        int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	UserName  string `json:"username,omitempty" column:"username"`
	File_id   int64  `json:"file_id" column:"user_file_perm_file_id"`
	Given_by  string `json:"given_by" column:"user_file_perm_given_by"`
	Perm_type int64  `json:"perm_type" column:"user_file_perm_type"`
	// Created_at string `json:"created_at" column:"user_file_perm_created_at"`
}

func (userfilepermission *UserFilePermission) Table() string {
	return "user_file_perm"
}

func (userfilepermission *UserFilePermission) String() string {
	return Stringify(userfilepermission)
}

type GroupFilePermission struct {
	Id        int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	Gr_Name   string `json:"gr_name,omitempty" column:"groupname"`
	File_id   int64  `json:"file_id" column:"grp_file_perm_file_id"`
	Given_by  string `json:"given_by" column:"grp_file_perm_given_by"`
	Perm_type int64  `json:"perm_type" column:"grp_file_perm_type"`
	// Created_at string `json:"created_at" column:"grp_file_perm_created_at"`
}

func (groupfilepermission *GroupFilePermission) Table() string {
	return "grp_file_perm"
}

func (groupfilepermission *GroupFilePermission) String() string {
	return Stringify(groupfilepermission)
}
