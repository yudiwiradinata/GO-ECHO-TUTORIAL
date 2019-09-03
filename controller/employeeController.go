package controller

import (
	"GO-ECHO-TUTORIAL/model"
	"GO-ECHO-TUTORIAL/repository"
	"GO-ECHO-TUTORIAL/service"
	"GO-ECHO-TUTORIAL/utility"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/dig"
)

var response = model.BaseResponse{}

//EmpController struct
type EmployeeController struct {
	EmpService service.Employee
	EmpRepo    repository.EmployeeRepo
}

type HParams struct {
	dig.In
	EmpService service.Employee
	EmpRepo    repository.EmployeeRepo
}

func NewEmployeeController(
	h HParams) *EmployeeController {
	return &EmployeeController{EmpService: h.EmpService,
		EmpRepo: h.EmpRepo}
}

//GetEmployees - get All employee
func (e EmployeeController) GetEmployees(c echo.Context) error {
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

//GetEmployeeByName - get All employee
func (e EmployeeController) GetEmployeeByName(c echo.Context) error {
	name := c.Param("name")

	result, err := e.EmpService.GetEmployeeByName(name, e.EmpRepo)
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
func (e EmployeeController) CreateEmployee(c echo.Context) error {
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
func (e EmployeeController) UpdateEmployee(c echo.Context) error {
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
func (e EmployeeController) DeleteEmployee(c echo.Context) error {
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
