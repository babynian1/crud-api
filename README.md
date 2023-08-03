Installation
- Clone repo
- Buat DB
- create table users  :
  create table users(
    -> id int not null auto_increment primary key
    -> , nama varchar(100),
    -> email varchar(100),
    -> umur int(100),
    -> created_at timestamp default current_timestamp,
    -> updated_at timestamp,
    -> deleted_at timestamp
    -> );
- set db .env
- go run main.go


API Documentation
https://www.postman.com/babynian/workspace/crud-api-golang/collection/8553128-13a7483d-249b-4945-80b5-9f73fd4446f1?action=share&creator=8553128
