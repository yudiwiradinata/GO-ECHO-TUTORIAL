package interfaces

import "GO-ECHO-TUTORIAL/model"

//EmployeeService - api
type EmployeeService interface {
	GetEmp(repo EmployeeRepository) (res model.Employees, err error)
	CreateEmp(empRequest model.EmployeeRequest, repo EmployeeRepository) (err error)
	DeleteEmp(name string, repo EmployeeRepository) (bool, error)
	UpdateEmp(id string, empRequest model.EmployeeRequest, repo EmployeeRepository) (bool, error)
}

//EmployeeRepository - api
type EmployeeRepository interface {
	GetEmployee() (res model.Employees, err error)
	InsertEmployee(req model.EmployeeRequest) (err error)
	UpdateEmployee(id string, req model.EmployeeRequest) (bool, error)
	DeleteEmployee(name string) (bool, error)
}
