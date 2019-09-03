package service

import (
	"GO-ECHO-TUTORIAL/model"

	"GO-ECHO-TUTORIAL/repository"
)

type employeeService struct {
	repository repository.EmployeeRepo
}

//Employee - interface
type Employee interface {
	GetEmp(repo repository.EmployeeRepo) (res model.Employees, err error)
	GetEmployeeByName(name string, repo repository.EmployeeRepo) (res model.Employees, err error)
	CreateEmp(empRequest model.EmployeeRequest, repo repository.EmployeeRepo) (err error)
	DeleteEmp(name string, repo repository.EmployeeRepo) (bool, error)
	UpdateEmp(id string, empRequest model.EmployeeRequest, repo repository.EmployeeRepo) (bool, error)
}

func NewEmployee(repository repository.EmployeeRepo) Employee {
	return &employeeService{
		repository: repository,
	}
}

func NewEmployeeService(repository repository.EmployeeRepo) *employeeService {
	return &employeeService{repository: repository}
}

//GetEmp - get
func (s employeeService) GetEmp(repo repository.EmployeeRepo) (res model.Employees, err error) {
	res, err = repo.GetEmployee()
	return res, err
}

//DeleteEmp - delete
func (s employeeService) DeleteEmp(name string, repo repository.EmployeeRepo) (bool, error) {
	res, err := repo.DeleteEmployee(name)
	return res, err
}

//UpdateEmp - update
func (s employeeService) UpdateEmp(id string, empRequest model.EmployeeRequest, repo repository.EmployeeRepo) (bool, error) {
	res, err := repo.UpdateEmployee(id, empRequest)
	return res, err
}

//CreateEmp - create
func (s employeeService) CreateEmp(empRequest model.EmployeeRequest, repo repository.EmployeeRepo) (err error) {
	err = repo.InsertEmployee(empRequest)
	return err
}

//GetEmployeeByName - get
func (s employeeService) GetEmployeeByName(name string, repo repository.EmployeeRepo) (res model.Employees, err error) {
	res, err = repo.GetEmployeeByName(name)
	return res, err
}
