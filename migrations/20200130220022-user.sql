
-- +migrate Up
CREATE TABLE user (
	id integer(11) AUTO_INCREMENT,
	auth_id integer(11) NOT NULL,
	Fam varchar(64) DEFAULT NULL,
	Name varchar(64) DEFAULT NULL,
	Otch varchar(64) DEFAULT NULL,
	Birthday date DEFAULT NULL,
	Description varchar(255) DEFAULT NULL,
	Avatar varchar(255) DEFAULT NULL,
	PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE user;