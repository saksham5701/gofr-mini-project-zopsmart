package handlers

import (
	"net/http"
	"strconv"
	"employee_manager/internal/models"
	"employee_manager/internal/service"
    "gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)


type EmployeeHandler struct {
	Service *service.EmployeeService
}


func (h *EmployeeHandler) GetEmployeesHandler(c *gofr.Context) (interface{}, error) {
	employees, err := h.Service.GetEmployees(c)
	if err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "There is some problem in the server"}
	}
return employees, nil
}
func (h *EmployeeHandler) AddEmployeeHandler(c *gofr.Context) (interface{}, error) {
	var employee models.Employee
	if err := c.Bind(&employee); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusBadRequest, Code: "BAD_REQUEST", Reason: "Wrong or invalid request"}
	}

	if err := h.Service.AddEmployee(c, employee); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "There is some problem in the server"}
	}

	return employee, nil
}


func (h *EmployeeHandler) UpdateEmployeeHandler(c *gofr.Context) (interface{}, error) {
	employeeID, err := strconv.Atoi(c.PathParam("eid"))
	if err != nil {
		return nil, &errors.Response{StatusCode: http.StatusBadRequest, Code: "BAD_REQUEST", Reason: "Invalid employee ID"}
	}

	var updatedEmployee models.Employee
	if err := c.Bind(&updatedEmployee); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusBadRequest, Code: "BAD_REQUEST", Reason: "Wrong or invalid request"}
	}

	updatedEmployee.Eid = employeeID;

	if err := h.Service.UpdateEmployee(c, updatedEmployee); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "There is some problem in the server"}
	}

	return updatedEmployee, nil
}


func (h *EmployeeHandler) DeleteEmployeeHandler(c *gofr.Context) (interface{}, error) {
	employeeID, err := strconv.Atoi(c.PathParam("eid"))
	if err != nil {
		return nil, &errors.Response{StatusCode: http.StatusBadRequest, Code: "BAD_REQUEST", Reason: "Invalid employee ID"}
	}

	
	if err := h.Service.DeleteEmployee(c, employeeID); err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "There is some problem in the server"}
	}

	
	employees, err := h.Service.GetEmployees(c)
	if err != nil {
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_SERVER_ERROR", Reason: "There is some problem in the server"}
	}

	return employees, nil
}
