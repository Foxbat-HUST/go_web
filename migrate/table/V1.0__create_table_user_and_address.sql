CREATE TABLE users(
id int PRIMARY KEY AUTO_INCREMENT NOT NULL,
name varchar(30),
age int,
email varchar(30),
created_at datetime DEFAULT CURRENT_TIMESTAMP,
updated_at datetime ON UPDATE CURRENT_TIMESTAMP,
deleted_at datetime DEFAULT NULL);
