
-- +migrate Up
CREATE TABLE user (
	id integer(11) AUTO_INCREMENT,
	auth_id integer(11) NOT NULL unique,
	fam varchar(64) DEFAULT NULL,
	name varchar(64) DEFAULT NULL,
	otch varchar(64) DEFAULT NULL,
	birthday date DEFAULT NULL,
	description varchar(255) DEFAULT NULL,
	avatar varchar(255) DEFAULT NULL,
	PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE user;