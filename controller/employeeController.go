package controller

import (
	"GO-ECHO-TUTORIAL/interfaces"
	"GO-ECHO-TUTORIAL/model"
	"GO-ECHO-TUTORIAL/utility"
	"net/http"

	"github.com/labstack/echo"
)

var response = model.BaseResponse{}

//EmpController struct
type EmpController struct {
	EmpService interfaces.EmployeeService
	EmpRepo    interfaces.EmployeeRepository
}

//GetEmployees - get All employee
func (e EmpController) GetEmployees(c echo.Context) error {
	result, err := e.EmpService.GetEmp(e.EmpRepo)
	if err != nil {
		response.Code = utility.CodeFailed
		response.Message = utility.ResponseFailed
		response.Data = err
	}

	response.Code = utility.CodeSuccess
	response.Message = utility.ResponseSuccess
	response.Data = result

	return c.JSON(http.StatusOK, response)
}

//CreateEmployee d
func (e EmpController) CreateEmployee(c echo.Context) error {
	emp := new(model.EmployeeRequest)
	if err := c.Bind(emp); err != nil {
		return err
	}

	err := e.EmpService.CreateEmp(*emp, e.EmpRepo)

	if err != nil {
		response.Code = utility.CodeFailed
		response.Message = utility.ResponseFailed
		response.Data = err
	}

	response.Code = utility.CodeSuccess
	response.Message = utility.ResponseSuccess
	response.Data = true

	return c.JSON(http.StatusOK, response)
}

//UpdateEmployee d
func (e EmpController) UpdateEmployee(c echo.Context) error {
	emp := new(model.EmployeeRequest)
	id := c.Param("id")
	if err := c.Bind(emp); err != nil {
		return err
	}
	res, err := e.EmpService.UpdateEmp(id, *emp, e.EmpRepo)

	if err != nil {
		response.Code = utility.CodeFailed
		response.Message = utility.ResponseFailed
		response.Data = res
	}

	response.Code = utility.CodeSuccess
	response.Message = utility.ResponseSuccess
	response.Data = res

	return c.JSON(http.StatusOK, response)
}

//DeleteEmployee d
func (e EmpController) DeleteEmployee(c echo.Context) error {
	name := c.Param("name")
	res, err := e.EmpService.DeleteEmp(name, e.EmpRepo)
	if err != nil {
		response.Code = utility.CodeFailed
		response.Message = utility.ResponseFailed
		response.Data = res
	}

	response.Code = utility.CodeSuccess
	response.Message = utility.ResponseSuccess
	response.Data = res

	return c.JSON(http.StatusOK, response)
}
