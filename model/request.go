package model

type EmployeeRequest struct {
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}
