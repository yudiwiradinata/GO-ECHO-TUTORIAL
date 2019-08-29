package service

import (
	"GO-ECHO-TUTORIAL/interfaces"
	"GO-ECHO-TUTORIAL/model"
)

type GetEmployee struct {
}

func (s GetEmployee) GetEmp(repo interfaces.EmployeeRepository) (res model.Employees, err error) {
	res, err = repo.GetEmployee()
	return res, err
}

func (s GetEmployee) DeleteEmp(name string, repo interfaces.EmployeeRepository) (bool, error) {
	res, err := repo.DeleteEmployee(name)
	return res, err
}

//UpdateEmp - update
func (s GetEmployee) UpdateEmp(id string, empRequest model.EmployeeRequest, repo interfaces.EmployeeRepository) (bool, error) {
	res, err := repo.UpdateEmployee(id, empRequest)
	return res, err
}

func (s GetEmployee) CreateEmp(empRequest model.EmployeeRequest, repo interfaces.EmployeeRepository) (err error) {
	err = repo.InsertEmployee(empRequest)
	return err
}
