
# Employee Management System API using Gofr

This an API project that helps you with Management of your employees records.It uses database as SQL.It performs CURD operations to do Creation,updation,Read and Deleting of a Record.


## To install Gofr


```bash
  go get gofr.dev
```
      go mod init employee_manager

      go mod tidy

      go run main.go

## Connection to MYSQL Database

```bash
  docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=employees_db -p 3306:3306 -d mysql:8.0.30

  docker exec -it gofr-mysql mysql -uroot -proot123 employees_db -e "CREATE TABLE employees (eid int(255) PRIMARY KEY, empname VARCHAR(255) NOT NULL,salary int(50),email VARCHAR(255));"
```
    
    
## Various API requests over Postman

### Get API
![GET](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/e386559c-afde-411b-adb5-6c771fbdad73)

### Post
![POST](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/908023a4-4545-4308-9bc1-21b2bad419c3)
### PUT
![PUT](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/272452b5-08b2-4522-b429-87f8c0a7f94b)
### DELETE
![DELETE](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/f1b4170e-16f0-46dd-af7c-cccea4cf078a)
## UML Diagram for the project
![UML](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/afd4512d-0b43-4b0d-8439-4e50f24cad7a)
 ## References
 - [Gofr Documentation ](https://gofr.dev/)
 - [Article over medium](https://medium.com/@mundhraumang.02/sample-rest-api-using-go-gorm-and-gofr-0ea41eaa6c62#:~:text=In%20the%20journey%20to%20create,server%20to%20handle%20incoming%20requests.&text=GoFr%20excels%20in%20simplifying%20the%20process%20of%20setting%20up%20a%20web%20server.) 








