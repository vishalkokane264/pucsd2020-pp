DROP DATABASE IF EXISTS restapi;
CREATE DATABASE IF NOT EXISTS restapi;
use restapi;
DROP TABLE IF EXISTS file_server;
DROP TABLE IF EXISTS group_file_perm;
DROP TABLE IF EXISTS user_file_perm;
DROP TABLE IF EXISTS user_group;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS groups;
-- DROP TABLE IF EXISTS permission;
DROP TABLE IF EXISTS user_role;

CREATE TABLE IF NOT EXISTS user_role(
    role_id int AUTO_INCREMENT PRIMARY KEY,
    role_name varchar(20) NOT NULL
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user_role values(1,'Admin');
INSERT INTO user_role values(2,'Read');
INSERT INTO user_role values(3,'read/write');

-- CREATE TABLE IF NOT EXISTS permission(
--     perm_id int PRIMARY KEY,
--     perm_type varchar(10)
-- )ENGINE = INNODB CHARACTER SET=utf8;

-- INSERT INTO permission values(1,'read');
-- INSERT INTO permission values(2,'write');
-- INSERT INTO permission values(3,'read/write');

CREATE TABLE IF NOT EXISTS user(
    user_id int AUTO_INCREMENT PRIMARY KEY,
    user_fname varchar(20) NOT NULL,
    user_lname varchar(20) NOT NULL,
    password varchar(20) NOT NULL
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user values(1,'fuser1','luser1','xyz');
INSERT INTO user values(2,'fuser2','luser2','xyz');
INSERT INTO user values(3,'fuser3','luser3','xyz');
INSERT INTO user values(4,'fuser4','luser4','xyz');
INSERT INTO user values(5,'fuser5','luser5','xyz');
INSERT INTO user values(6,'fuser6','luser6','xyz');
INSERT INTO user values(7,'fuser7','luser7','xyz');
INSERT INTO user values(8,'fuser8','luser8','xyz');
INSERT INTO user values(9,'fuser9','luser9','xyz');
INSERT INTO user values(10,'fuser10','luser10','xyz');
INSERT INTO user values(11,'fuser11','luser11','xyz');

/*New table created*/
CREATE TABLE IF NOT EXISTS user_default_perm(
    user_default_perm_user_id int NOT NULL,
    user_default_perm_user_type_id int NOT NULL,
    -- user_default_perm int NOT NULL,
    FOREIGN KEY(user_default_perm_user_id) REFERENCES user(user_id),
    FOREIGN KEY(user_default_perm_user_type_id) REFERENCES user_role(role_id),
    -- FOREIGN KEY(user_default_perm) REFERENCES permission(perm_id),
    PRIMARY KEY(user_default_perm_user_id,user_default_perm_user_type_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user_default_perm values(1,1);
INSERT INTO user_default_perm values(1,2);
INSERT INTO user_default_perm values(2,2);
INSERT INTO user_default_perm values(3,2);
INSERT INTO user_default_perm values(4,1);
INSERT INTO user_default_perm values(4,2);
INSERT INTO user_default_perm values(5,1);
INSERT INTO user_default_perm values(5,2);
INSERT INTO user_default_perm values(6,3);
INSERT INTO user_default_perm values(7,2);
INSERT INTO user_default_perm values(8,2);
INSERT INTO user_default_perm values(9,2);
INSERT INTO user_default_perm values(10,2);
INSERT INTO user_default_perm values(11,2);


CREATE TABLE IF NOT EXISTS groups(
    gr_id int AUTO_INCREMENT NOT NULL,
    gr_name varchar(20) NOT NULL,
    gr_owner int NOT NULL,
    FOREIGN KEY(gr_owner) REFERENCES user_default_perm(user_default_perm_user_id),
    PRIMARY KEY(gr_id,gr_owner)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO groups values(1,'grp1',1);
INSERT INTO groups values(2,'grp2',1);
INSERT INTO groups values(3,'grp3',5);
INSERT INTO groups values(4,'grp4',5);

CREATE TABLE IF NOT EXISTS file_server(
    file_id int AUTO_INCREMENT,
    file_path varchar(255) NOT NULL,
    owner_id int NOT NULL,
    FOREIGN KEY(owner_id) REFERENCES user_default_perm(user_default_perm_user_id), 
    PRIMARY KEY(file_id,owner_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO file_server values(1,'/home/vishal/',1);
INSERT INTO file_server values(2,'/home/vishal/Documents',1);
INSERT INTO file_server values(3,'/home/vishal/Documents/Mydata',1);
INSERT INTO file_server values(4,'/home/vishal/Pictures',1);
INSERT INTO file_server values(5,'/home/vishal/Pictures/image.jpg',1);
INSERT INTO file_server values(6,'/home/postgres',5);
INSERT INTO file_server values(7,'/home/postgres/Music',5);
INSERT INTO file_server values(8,'/home/postgres/Videos',5);

/*file_mapping table: for checking which file is a subfile of other */
CREATE TABLE IF NOT EXISTS file_mapping(
    file_mapping_file_id int NOT NULL,
    file_mapping_file_subfile_id int DEFAULT NULL,
    FOREIGN KEY (file_mapping_file_id) REFERENCES file_server(file_id),
    FOREIGN KEY (file_mapping_file_subfile_id) REFERENCES file_server(file_id),
    UNIQUE (file_mapping_file_id,file_mapping_file_subfile_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO file_mapping(file_mapping_file_id,file_mapping_file_subfile_id) VALUES(1,2);
INSERT INTO file_mapping(file_mapping_file_id,file_mapping_file_subfile_id) VALUES(1,3);
INSERT INTO file_mapping(file_mapping_file_id,file_mapping_file_subfile_id) VALUES(1,4);
INSERT INTO file_mapping(file_mapping_file_id,file_mapping_file_subfile_id) VALUES(1,5);
INSERT INTO file_mapping(file_mapping_file_id,file_mapping_file_subfile_id) VALUES(2,3);
INSERT INTO file_mapping(file_mapping_file_id) VALUES(3);
INSERT INTO file_mapping(file_mapping_file_id,file_mapping_file_subfile_id) VALUES(4,5);
INSERT INTO file_mapping(file_mapping_file_id) VALUES(5);
INSERT INTO file_mapping(file_mapping_file_id,file_mapping_file_subfile_id) VALUES(6,7);
INSERT INTO file_mapping(file_mapping_file_id,file_mapping_file_subfile_id) VALUES(6,8);
INSERT INTO file_mapping(file_mapping_file_id) VALUES(7);
INSERT INTO file_mapping(file_mapping_file_id) VALUES(8);


CREATE TABLE IF NOT EXISTS user_group(
    user_grp_user_id int NOT NULL,  
    user_grp_gr_id int NOT NULL,
    FOREIGN KEY(user_grp_gr_id) REFERENCES groups(gr_id),
    FOREIGN KEY(user_grp_user_id) REFERENCES user_default_perm(user_default_perm_user_id),
    PRIMARY KEY(user_grp_user_id,user_grp_gr_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user_group values(2,1);
INSERT INTO user_group values(3,1);
INSERT INTO user_group values(6,1);
INSERT INTO user_group values(7,1);
INSERT INTO user_group values(8,2);
INSERT INTO user_group values(9,2);
INSERT INTO user_group values(10,2);
INSERT INTO user_group values(11,2);
INSERT INTO user_group values(2,3);
INSERT INTO user_group values(3,3);
INSERT INTO user_group values(10,3);
INSERT INTO user_group values(11,3);
INSERT INTO user_group values(6,4);
INSERT INTO user_group values(7,4);
INSERT INTO user_group values(8,4);
INSERT INTO user_group values(9,4);


CREATE TABLE IF NOT EXISTS user_file_perm(
    user_file_perm_user_id int NOT NULL,
    user_file_perm_file_id int NOT NULL,
    user_file_perm_given_by int NOT NULL,
    user_file_perm_type int NOT NULL,
    user_file_perm_created_at datetime DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_file_perm_user_id) REFERENCES user_default_perm(user_default_perm_user_id),
    FOREIGN KEY(user_file_perm_given_by) REFERENCES user_default_perm(user_default_perm_user_id),
    FOREIGN KEY(user_file_perm_file_id) REFERENCES file_server(file_id),
    FOREIGN KEY(user_file_perm_type) REFERENCES user_role(role_id),
    PRIMARY KEY(user_file_perm_user_id,user_file_perm_file_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO user_file_perm values(2,2,1,3,CURRENT_TIMESTAMP);
INSERT INTO user_file_perm values(6,6,5,3,CURRENT_TIMESTAMP);
INSERT INTO user_file_perm values(2,5,1,1,CURRENT_TIMESTAMP);
INSERT INTO user_file_perm values(3,7,5,3,CURRENT_TIMESTAMP);


CREATE TABLE IF NOT EXISTS grp_file_perm(
    grp_file_perm_grp_id int NOT NULL,
    grp_file_perm_file_id int NOT NULL,
    grp_file_perm_type int NOT NULL,
    grp_file_perm_given_by int NOT NULL,
    grp_file_perm_created_at datetime DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(grp_file_perm_grp_id) REFERENCES groups(gr_id),
    FOREIGN KEY(grp_file_perm_given_by) REFERENCES user_default_perm(user_default_perm_user_id),
    FOREIGN KEY(grp_file_perm_file_id) REFERENCES file_server(file_id),
    FOREIGN KEY(grp_file_perm_type) REFERENCES user_role(role_id),
    PRIMARY KEY(grp_file_perm_grp_id,grp_file_perm_file_id)
)ENGINE = INNODB CHARACTER SET=utf8;

INSERT INTO grp_file_perm values(1,1,3,1,CURRENT_TIMESTAMP);
INSERT INTO grp_file_perm values(2,6,3,5,CURRENT_TIMESTAMP);
INSERT INTO grp_file_perm values(4,5,3,5,CURRENT_TIMESTAMP);
INSERT INTO grp_file_perm values(3,6,3,5,CURRENT_TIMESTAMP);
INSERT INTO grp_file_perm values(4,8,3,5,CURRENT_TIMESTAMP);

