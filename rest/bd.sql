-- para crear bd primero loggin
mysql -u root -p
-- crear bd
create database goweb;
-- seleccionar base de datos
use goweb

CREATE TABLE users(
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(64) NOT NULL,
    email VARCHAR(40),
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP );



-- ver tablas
show tables;
show columns from users;