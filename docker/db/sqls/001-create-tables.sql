create table users
(id int not null auto_increment primary key, user_id varchar(100) not null unique, user_name varchar(100),price int)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table customers
(id int not null auto_increment primary key, customer_id varchar(100) not null unique,password varchar(20), gender varchar(100),customer_name varchar(50),address varchar(100),birth datetime, role varchar(20))DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table stores
(id int not null auto_increment primary key, store_id varchar(100) not null unique,password varchar(20), store_name varchar(100),category varchar(50),latitude float ,longitude float,address varchar(100),price int)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table terms
(id int not null auto_increment primary key, store_id varchar(100) not null unique,gender varchar(10), min_age int,max_age int,role varchar(20) )DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

INSERT INTO users (user_id, user_name,price) VALUES ("id","hogeo",100);
INSERT INTO users (user_id, user_name,price) VALUES ("id2","hogehoge2",800);