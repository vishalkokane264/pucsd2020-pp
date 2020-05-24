DROP DATABASE IF EXISTS restapi;
CREATE DATABASE IF NOT EXISTS restapi;
use restapi;
CREATE TABLE IF NOT EXISTS user(
    id int AUTO_INCREMENT PRIMARY KEY,
    user_fname varchar(20) NOT NULL,
    user_lname varchar(20) NOT NULL,
    username varchar(20) NOT NULL, 
    password varchar(20) NOT NULL,
    UNIQUE(username)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user values(1,'fname','lname','adminuser','password');

CREATE TABLE IF NOT EXISTS user_role(
    id int AUTO_INCREMENT PRIMARY KEY,
    role_name varchar(20) NOT NULL UNIQUE
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user_role values(1,'Admin');
INSERT INTO user_role values(2,'Read');
INSERT INTO user_role values(3,'Read/write');

CREATE TABLE IF NOT EXISTS user_default_perm(
    id int AUTO_INCREMENT PRIMARY KEY,
    username varchar(20) NOT NULL,
    user_default_perm_user_type_id int,
    CONSTRAINT fk_uname FOREIGN KEY(username) REFERENCES user(username) on delete cascade,
    UNIQUE(username,user_default_perm_user_type_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user_default_perm values(1,'adminuser',1);


CREATE TABLE IF NOT EXISTS groups(
    id int AUTO_INCREMENT PRIMARY KEY,
    gr_name varchar(20) NOT NULL,
    gr_owner varchar(20) NOT NULL,
    CONSTRAINT fk_gowner FOREIGN KEY(gr_owner) REFERENCES user_default_perm(username) on delete cascade,
    UNIQUE(gr_name,gr_owner)
)ENGINE = INNODB CHARACTER SET=utf8;

/*INSERT INTO groups values(1,'grp1','adminuser');*/

CREATE TABLE IF NOT EXISTS user_group(
    username varchar(20) NOT NULL,  
    user_grp_gr_name varchar(20) NOT NULL,
    FOREIGN KEY(username) REFERENCES user_default_perm(username) on delete cascade,
    FOREIGN KEY(user_grp_gr_name) REFERENCES groups(gr_name) on delete cascade,
    PRIMARY KEY(username,user_grp_gr_name)
)ENGINE = INNODB CHARACTER SET=utf8;


CREATE TABLE IF NOT EXISTS file_server(
    id int AUTO_INCREMENT,
    file_path varchar(255) NOT NULL,
    owner_name varchar(20) NOT NULL,
    FOREIGN KEY(owner_name) REFERENCES user_default_perm(username), 
    PRIMARY KEY(id,owner_name)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO file_server values(1,'/home/vishal/','adminuser');
INSERT INTO file_server values(2,'/home/vishal/Documents','adminuser');
INSERT INTO file_server values(3,'/home/vishal/Documents/Mydata','adminuser');
INSERT INTO file_server values(4,'/home/vishal/Pictures','adminuser');
INSERT INTO file_server values(5,'/home/vishal/Pictures/image.jpg','adminuser');
INSERT INTO file_server values(6,'/home/postgres','adminuser');
INSERT INTO file_server values(7,'/home/postgres/Music','adminuser');
INSERT INTO file_server values(8,'/home/postgres/Videos','adminuser');

CREATE TABLE IF NOT EXISTS user_file_perm(
    id int AUTO_INCREMENT PRIMARY KEY,
    username varchar(20) NOT NULL,
    user_file_perm_file_id int NOT NULL,
    user_file_perm_given_by varchar(20) NOT NULL,
    user_file_perm_type int NOT NULL CHECK (user_file_perm_type IN (2,3)),
    FOREIGN KEY(username) REFERENCES user_default_perm(username) on delete cascade,
    FOREIGN KEY(user_file_perm_given_by) REFERENCES user_default_perm(username) on delete cascade,
    FOREIGN KEY(user_file_perm_file_id) REFERENCES file_server(id) on delete cascade,
    UNIQUE(username,user_file_perm_file_id)
)ENGINE = INNODB CHARACTER SET=utf8;

CREATE TABLE IF NOT EXISTS grp_file_perm(
    id int AUTO_INCREMENT PRIMARY KEY,
    groupname varchar(20) NOT NULL,
    grp_file_perm_file_id int NOT NULL,
    grp_file_perm_given_by varchar(20) NOT NULL,
    grp_file_perm_type int NOT NULL CHECK (user_file_perm_type IN (2,3)),
    FOREIGN KEY(groupname) REFERENCES groups(gr_name) on delete cascade,
    FOREIGN KEY(grp_file_perm_given_by) REFERENCES user_default_perm(username) on delete cascade,
    FOREIGN KEY(grp_file_perm_file_id) REFERENCES file_server(id) on delete cascade,
    UNIQUE(groupname,grp_file_perm_file_id)
)ENGINE = INNODB CHARACTER SET=utf8;
