create database am;

use am;

create table if not exists user (
    username varchar(32) primary key,
    password varchar(32) not null
);