package service

import (
	"GO-ECHO-TUTORIAL/model"

	"GO-ECHO-TUTORIAL/repository"
)

type employeeService struct {
	repository repository.EmployeeRepositoryInterface
}

//EmployeeServiceInterface - interface
type EmployeeServiceInterface interface {
	GetEmp() (res *model.Employees, err error)
	GetEmployeeByName(name string) (res *model.Employees, err error)
	CreateEmp(empRequest model.EmployeeRequest) (err error)
	DeleteEmp(name string) (bool, error)
	UpdateEmp(id string, empRequest model.EmployeeRequest) (bool, error)
}

func NewEmployeeServiceInterface(repository repository.EmployeeRepositoryInterface) EmployeeServiceInterface {
	return &employeeService{
		repository: repository,
	}
}

func NewEmployeeService(repository repository.EmployeeRepositoryInterface) *employeeService {
	return &employeeService{repository: repository}
}

//GetEmp - get
func (s employeeService) GetEmp() (res *model.Employees, err error) {
	res, err = s.repository.GetEmployee()
	return res, err
}

//DeleteEmp - delete
func (s employeeService) DeleteEmp(name string) (bool, error) {
	res, err := s.repository.DeleteEmployee(name)
	return res, err
}

//UpdateEmp - update
func (s employeeService) UpdateEmp(id string, empRequest model.EmployeeRequest) (bool, error) {
	res, err := s.repository.UpdateEmployee(id, empRequest)
	return res, err
}

//CreateEmp - create
func (s employeeService) CreateEmp(empRequest model.EmployeeRequest) (err error) {
	err = s.repository.InsertEmployee(empRequest)
	return err
}

//GetEmployeeByName - get
func (s employeeService) GetEmployeeByName(name string) (res *model.Employees, err error) {
	res, err = s.repository.GetEmployeeByName(name)
	return res, err
}
