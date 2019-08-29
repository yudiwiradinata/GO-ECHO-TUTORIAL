package repository

import (
	"GO-ECHO-TUTORIAL/db"
	"GO-ECHO-TUTORIAL/model"
	"database/sql"
	"fmt"
)

type GetEmployeeRepo struct{}

var con *sql.DB

//DeleteEmployee - delete
func (e GetEmployeeRepo) DeleteEmployee(name string) (bool, error) {
	con = db.CreateConn()
	sqlStatement := "DELETE FROM employee WHERE employee_name = ?"

	res, err := con.Exec(sqlStatement, name)
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

func (e GetEmployeeRepo) UpdateEmployee(id string, req model.EmployeeRequest) (bool, error) {
	con = db.CreateConn()
	sqlStatement := "UPDATE employee SET employee_name = ?,employee_age = ?, employee_salary=? WHERE id=?"

	res, err := con.Exec(sqlStatement, req.Name, req.Age, req.Salary, id)

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

func (e GetEmployeeRepo) InsertEmployee(req model.EmployeeRequest) (err error) {
	con = db.CreateConn()
	sqlStatement := "INSERT INTO employee (employee_name, employee_age, employee_salary) VALUES (?,?,?)"
	_, err = con.Exec(sqlStatement, req.Name, req.Age, req.Salary)
	if err != nil {
		return err
	}

	return err
}

func (e GetEmployeeRepo) GetEmployee() (res model.Employees, err error) {
	//db.CreateCon()
	con = db.CreateConn()

	sqlStatement := "SELECT id,employee_name, employee_age, employee_salary FROM employee order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
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

	return result, err

}
