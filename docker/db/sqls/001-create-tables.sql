create table students
(id int not null auto_increment primary key, password varchar(128) not null, userid varchar(100) not null unique, name varchar(100))DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

INSERT INTO students (password,userid, name) VALUES ("pass","id","hogeo");