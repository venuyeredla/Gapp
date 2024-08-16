mysql -uroot -pecompwd#24

show databases;
CREATE DATABASE ecom;
use ecom;
create table ecom_user(id int NOT NULL AUTO_INCREMENT, first_name varchar(25), last_name varchar(25), email varchar(30), pwd varchar(25), PRIMARY KEY (id))

ALTER TABLE ecom_user RENAME COLUMN eamil to email;

insert into ecom_user(first_name, last_name, email,pwd) values ("venu","gopal","venugopal@ecom.com","ecom#24");