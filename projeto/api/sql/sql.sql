show databases;
create database if not exists golang;
use golang;
drop table users;
desc users;


create table users(

    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(20) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;



select * from users