package model

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"employee_name"`
	Salary string `json:"employee_salary"`
	Age    string `json:"employee_age"`
}

type Employees struct {
	Employees []Employee `json:"employee"`
}
