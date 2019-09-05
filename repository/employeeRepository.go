package repository

import (
	"GO-ECHO-TUTORIAL/model"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type employeeRepository struct {
	database *sql.DB
}

//EmployeeRepositoryInterface - api
type EmployeeRepositoryInterface interface {
	GetEmployee() (res *model.Employees, err error)
	GetEmployeeByName(name string) (res *model.Employees, err error)
	GetEmployeeByID(id int64) (res *model.Employee, err error)
	InsertEmployee(req model.EmployeeRequest) (res *model.Employee, err error)
	UpdateEmployee(id string, req model.EmployeeRequest) (bool, error)
	DeleteEmployee(name string) (bool, error)
}

//NewEmployeeRepositoryInterface -
func NewEmployeeRepositoryInterface(con *sql.DB) EmployeeRepositoryInterface {
	return &employeeRepository{
		database: con,
	}
}

//DeleteEmployee - delete
func (e employeeRepository) DeleteEmployee(name string) (bool, error) {
	sqlStatement := sq.Delete("employee").Where("employee_name = ?", name)
	sql, args, err := sqlStatement.ToSql()
	res, err := e.database.Exec(sql, args[0])

	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err == nil {
		if count > 0 {
			return true, err
		}
	}
	return false, err
}

//UpdateEmployee - update
func (e employeeRepository) UpdateEmployee(id string, req model.EmployeeRequest) (bool, error) {
	//con = db.CreateConn()
	//sqlStatement := "UPDATE employee SET employee_name = ?,employee_age = ?, employee_salary=? WHERE id=?"
	//res, err := con.Exec(sqlStatement, req.Name, req.Age, req.Salary, id)
	sqlStatement := sq.Update("employee").Set("employee_name", req.Name).Set("employee_age", req.Age).Set("employee_salary", req.Salary).Where("id= ?", id)
	sql, args, err := sqlStatement.ToSql()
	res, err := e.database.Exec(sql, args[0], args[1], args[2], args[3])
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err == nil {
		if count > 0 {
			return true, err
		}
	}

	return false, err
}

//GetEmployeeById - emp by id
func (e employeeRepository) GetEmployeeByID(id int64) (res *model.Employee, err error) {
	//con = db.CreateConn()
	sqlStatement := sq.Select("id", "employee_name", "employee_age", "employee_salary").From("employee").
		Where("id = ?", id).OrderBy("id")
	rows, err := sqlStatement.RunWith(e.database).Query()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	employee := model.Employee{}
	for rows.Next() {
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Age, &employee.Salary)
		if err2 != nil {
			fmt.Print(err2)
		}
	}
	return &employee, err
}

//InsertEmployee - insert
func (e employeeRepository) InsertEmployee(req model.EmployeeRequest) (res *model.Employee, err error) {
	//con = db.CreateConn()
	//sqlStatement := "INSERT INTO employee (employee_name, employee_age, employee_salary) VALUES (?,?,?)"
	//_, err = con.Exec(sqlStatement, req.Name, req.Age, req.Salary)
	sqlStatement := sq.Insert("employee").Columns("employee_name", "employee_age", "employee_salary").Values(
		req.Name, req.Age, req.Salary,
	).RunWith(e.database)

	id, err := sqlStatement.Exec()
	if err != nil {
		return nil, err
	}
	idd, _ := id.LastInsertId()
	return e.GetEmployeeByID(idd)
}

//GetEmployee - getAll
func (e employeeRepository) GetEmployee() (res *model.Employees, err error) {
	//con = db.CreateConn()
	//sqlStatement := "SELECT id,employee_name, employee_age, employee_salary FROM employee order by id"
	//rows, err := con.Query(sqlStatement)
	sqlStatement := sq.Select("id", "employee_name", "employee_age", "employee_salary").From("employee").OrderBy("id")
	rows, err := sqlStatement.RunWith(e.database).Query()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := model.Employees{}
	for rows.Next() {
		employee := model.Employee{}
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Age, &employee.Salary)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return &result, err
}

//GetEmployeeByName - emp by name
func (e employeeRepository) GetEmployeeByName(name string) (res *model.Employees, err error) {
	//con = db.CreateConn()
	sqlStatement := sq.Select("id", "employee_name", "employee_age", "employee_salary").From("employee").
		Where("employee_name = ?", name).OrderBy("id")
	rows, err := sqlStatement.RunWith(e.database).Query()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := model.Employees{}
	for rows.Next() {
		employee := model.Employee{}
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Age, &employee.Salary)
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return &result, err
}
