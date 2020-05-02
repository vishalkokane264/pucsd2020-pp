CREATE DATABASE IF NOT EXISTS rbac;
use rbac;
DROP TABLE IF EXISTS group_file_perm;
DROP TABLE IF EXISTS user_file_perm;
DROP TABLE IF EXISTS file_server;
DROP TABLE IF EXISTS user_group;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS permission;
DROP TABLE IF EXISTS user_type;

CREATE TABLE IF NOT EXISTS user_type(
    type_id int PRIMARY KEY,
    type_name varchar(20) NOT NULL,
    type_desc text
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user_type values(1,'Admin',' user type Admin');
INSERT INTO user_type values(2,'User',' normal user');

CREATE TABLE IF NOT EXISTS permission(
    perm_id int PRIMARY KEY,
    value varchar(10)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO permission values(0,'false');
INSERT INTO permission values(1,'true');

CREATE TABLE IF NOT EXISTS groups(
    gr_id int AUTO_INCREMENT PRIMARY KEY,
    gr_name varchar(20) NOT NULL
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO groups values(1,'group1');
INSERT INTO groups values(2,'group2');
INSERT INTO groups values(3,'group3');

CREATE TABLE IF NOT EXISTS user(
    user_id int AUTO_INCREMENT PRIMARY KEY,
    user_name varchar(20) NOT NULL,
    user_type_id int NOT NULL,
    password varchar(20) NOT NULL,
    FOREIGN KEY(user_type_id) REFERENCES user_type(type_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user values(1,'admin_user',1,'1234');
INSERT INTO user values(2,'user2',2,'1234');
INSERT INTO user values(3,'user3',2,'1234');
INSERT INTO user values(4,'user4',2,'1234');
INSERT INTO user values(5,'user5',2,'1234');


CREATE TABLE IF NOT EXISTS file_server(
    file_id int PRIMARY KEY,
    path text NOT NULL,
    perm_read int NOT NULL,
    perm_write int NOT NULL,
    perm_exec int NOT NULL,
    FOREIGN KEY(perm_read) REFERENCES permission(perm_id),
    FOREIGN KEY(perm_write) REFERENCES permission(perm_id),
    FOREIGN KEY(perm_exec) REFERENCES permission(perm_id)
)ENGINE = INNODB CHARACTER SET=utf8;

CREATE TABLE IF NOT EXISTS user_group(
    user_grp_user_id int NOT NULL,
    user_grp_gr_id int NOT NULL,
    FOREIGN KEY(user_grp_gr_id) REFERENCES groups(gr_id),
    FOREIGN KEY(user_grp_user_id) REFERENCES user(user_id),
    PRIMARY KEY(user_grp_user_id,user_grp_gr_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user_group values(2,1);
INSERT INTO user_group values(3,1);
INSERT INTO user_group values(4,1);
INSERT INTO user_group values(4,2);
INSERT INTO user_group values(5,3);


CREATE TABLE IF NOT EXISTS user_file_perm(
    user_file_perm_user_id int NOT NULL,
    user_file_perm_file_id int NOT NULL,
    user_file_perm_read int NOT NULL,
    user_file_perm_write int NOT NULL,
    user_file_perm_exec int NOT NULL,
    grp_file_perm_createdat date NOT NULL,
    FOREIGN KEY(user_file_perm_user_id) REFERENCES user(user_id),
    FOREIGN KEY(user_file_perm_file_id) REFERENCES file_server(file_id),
    FOREIGN KEY(user_file_perm_read) REFERENCES permission(perm_id),
    FOREIGN KEY(user_file_perm_write) REFERENCES permission(perm_id),
    FOREIGN KEY(user_file_perm_exec) REFERENCES permission(perm_id),
    PRIMARY KEY(user_file_perm_user_id,user_file_perm_file_id)
)ENGINE = INNODB CHARACTER SET=utf8;

CREATE TABLE IF NOT EXISTS grp_file_perm(
    grp_file_perm_grp_id int NOT NULL,
    grp_file_perm_file_id int NOT NULL,
    grp_file_perm_read int  NOT NULL,
    grp_file_perm_write int NOT NULL,
    grp_file_perm_exec int NOT NULL,
    grp_file_perm_createdat date NOT NULL,
    FOREIGN KEY(grp_file_perm_grp_id) REFERENCES groups(gr_id),
    FOREIGN KEY(grp_file_perm_file_id) REFERENCES file_server(file_id),
    FOREIGN KEY(grp_file_perm_read) REFERENCES permission(perm_id),
    FOREIGN KEY(grp_file_perm_write) REFERENCES permission(perm_id),
    FOREIGN KEY(grp_file_perm_exec) REFERENCES permission(perm_id),
    PRIMARY KEY(grp_file_perm_grp_id,grp_file_perm_file_id)
)ENGINE = INNODB CHARACTER SET=utf8;

