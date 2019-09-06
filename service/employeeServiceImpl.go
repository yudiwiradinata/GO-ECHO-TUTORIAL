package service

import (
	"GO-ECHO-TUTORIAL/model"
	"GO-ECHO-TUTORIAL/utility"

	"GO-ECHO-TUTORIAL/repository"

	"go.uber.org/dig"
)

type employeeService struct {
	repository repository.EmployeeRepositoryInterface
	cache      CacheServiceInterface
}

type EmpServiceParams struct {
	dig.In

	Repository repository.EmployeeRepositoryInterface
	Cache      CacheServiceInterface
}

//EmployeeServiceInterface - interface
type EmployeeServiceInterface interface {
	GetEmp() (res *model.Employees, err error)
	GetEmployeeByName(name string) (res *model.Employees, err error)
	CreateEmp(empRequest model.EmployeeRequest) (err error)
	DeleteEmp(name string) (bool, error)
	UpdateEmp(id string, empRequest model.EmployeeRequest) (bool, error)
}

//NewEmployeeServiceInterface -
func NewEmployeeServiceInterface(h EmpServiceParams) EmployeeServiceInterface {
	return &employeeService{
		repository: h.Repository,
		cache:      h.Cache,
	}
}

//GetEmp - get
func (s employeeService) GetEmp() (res *model.Employees, err error) {
	res, err = s.repository.GetEmployee()
	return res, err
}

//DeleteEmp - delete
func (s employeeService) DeleteEmp(name string) (bool, error) {
	s.cache.DeleteCache(utility.CreateKeyRedis(name))
	res, err := s.repository.DeleteEmployee(name)
	return res, err
}

//UpdateEmp - update
func (s employeeService) UpdateEmp(id string, empRequest model.EmployeeRequest) (bool, error) {
	s.cache.DeleteCache(utility.CreateKeyRedis(empRequest.Name))
	res, err := s.repository.UpdateEmployee(id, empRequest)
	return res, err
}

//CreateEmp - create
func (s employeeService) CreateEmp(empRequest model.EmployeeRequest) error {
	emp, err := s.repository.InsertEmployee(empRequest)
	if err != nil {
		return err
	}
	return s.cache.CreateCache(utility.CreateKeyRedis(empRequest.Name), *emp, 0)
}

//GetEmployeeByName - get
func (s employeeService) GetEmployeeByName(name string) (res *model.Employees, err error) {
	key := utility.CreateKeyRedis(name)
	var emp model.Employee
	err = s.cache.FindCache(key, &emp)
	if err != nil {
		res, err := s.repository.GetEmployeeByName(name)
		if res != nil {
			s.cache.CreateCache(utility.CreateKeyRedis(name), res.Employees[0], 0)
		}
		return res, err
	}
	var empS []model.Employee
	empS = append(empS, emp)
	res2 := model.Employees{
		Employees: empS,
	}
	return &res2, err
}
