
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
  docker run --name employee-project-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=employee_db -p 3306:3306 -d mysql:8.0.30
```
    
    
## Various API requests over Postman

### Get API

![Read](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/69695020-b767-4cd5-9057-1410f9ebaf4a)

### Post
![Create](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/ea3da51a-2a81-47f8-9fd9-ef8e5d9a05f9)
### PUT
![Update](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/d2b4e0d0-2cdf-4e00-ac33-dd0cde6f4c0d)
### DELETE
![delete](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/0f29c45b-a2fa-41b7-a970-9949d59cad8a)

## UML Diagram for the project
![UML Diagram](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/6b576a4e-a7ab-4621-8448-d35ad2b21869)

 ## References
 - [Gofr Documentation ](https://gofr.dev/)
 - [Article over medium](https://medium.com/@mundhraumang.02/sample-rest-api-using-go-gorm-and-gofr-0ea41eaa6c62#:~:text=In%20the%20journey%20to%20create,server%20to%20handle%20incoming%20requests.&text=GoFr%20excels%20in%20simplifying%20the%20process%20of%20setting%20up%20a%20web%20server.) 


