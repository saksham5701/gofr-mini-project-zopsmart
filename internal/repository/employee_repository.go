package repository

import (
	"database/sql"
	"employee_manager/internal/models"
     "gofr.dev/pkg/gofr"
)


type EmployeeRepository struct {
	DB *sql.DB
}


func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}


func (r *EmployeeRepository) GetEmployees(ctx *gofr.Context) ([]models.Employee, error) {
	rows, err := ctx.DB().QueryContext(ctx, "SELECT eid, empname, salary, email FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.Eid,&employee.Empname, &employee.Salary, &employee.Email); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}


func (r *EmployeeRepository) AddEmployee(ctx *gofr.Context, employee models.Employee) error {
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO employees (eid,empname, salary, email) VALUES (?,?, ?, ?)",employee.Eid, employee.Empname, employee.Salary, employee.Email)
	return err
}


func (r *EmployeeRepository) UpdateEmployee(ctx *gofr.Context, employee models.Employee) error {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE employees SET empname=?,salary=?, email=? WHERE eid=?", employee.Empname,employee.Salary, employee.Email,employee.Eid)
	return err
}


func (r *EmployeeRepository) DeleteEmployee(ctx *gofr.Context, employeeID int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM employees WHERE eid=?", employeeID)
	return err
}
