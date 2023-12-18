package service

import (
	"employee_manager/internal/models"
	"employee_manager/internal/repository"
    "gofr.dev/pkg/gofr"
)


type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}
func (s *EmployeeService) GetEmployees(ctx *gofr.Context) ([]models.Employee, error) {
	return s.Repo.GetEmployees(ctx)
}
func (s *EmployeeService) AddEmployee(ctx *gofr.Context, employee models.Employee) error {
	return s.Repo.AddEmployee(ctx, employee)
}
func (s *EmployeeService) UpdateEmployee(ctx *gofr.Context, employee models.Employee) error {
	return s.Repo.UpdateEmployee(ctx, employee)
}
func (s *EmployeeService) DeleteEmployee(ctx *gofr.Context, employeeID int) error {
	return s.Repo.DeleteEmployee(ctx, employeeID)
}
