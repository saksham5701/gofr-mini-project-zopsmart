package main

import (
	"database/sql"
	"log"

	"gofr.dev/pkg/gofr"

	"employee_manager/internal/handlers"
	"employee_manager/internal/repository"
	"employee_manager/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/employees_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	EmployeeRepo := repository.NewEmployeeRepository(db)
	EmployeeService := service.NewEmployeeService(EmployeeRepo)
	EmployeeHandler := &handlers.EmployeeHandler{Service: EmployeeService}

	app := gofr.New()
    app.GET("/employees", EmployeeHandler.GetEmployeesHandler)
	app.POST("/employee", EmployeeHandler.AddEmployeeHandler)
	app.PUT("/employee/eid", EmployeeHandler.UpdateEmployeeHandler)
	app.DELETE("/employee/eid", EmployeeHandler.DeleteEmployeeHandler)

	app.Start()
}

