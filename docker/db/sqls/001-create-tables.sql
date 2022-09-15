create table users
(id int not null auto_increment primary key, user_id varchar(100) not null unique, user_name varchar(100),price int)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;



INSERT INTO users (user_id, user_name,price) VALUES ("id","hogeo",100);
INSERT INTO users (user_id, user_name,price) VALUES ("id2","hogehoge2",800);