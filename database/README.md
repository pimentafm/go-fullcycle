`docker compose up -d`

`docker exec -it mysql bash`

`mysql -h localhost -P 3306 -u root -p`

`create database goexpert;`

`use goexpert;`

`create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));`


```go
go clean -modcache
go mod tidy
```
