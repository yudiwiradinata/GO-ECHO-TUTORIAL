package main

import (
	"GO-ECHO-TUTORIAL/controller"
	"GO-ECHO-TUTORIAL/repository"
	"GO-ECHO-TUTORIAL/service"
	"GO-ECHO-TUTORIAL/utility"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	c := controller.EmpController{
		EmpService: service.GetEmployee{},
		EmpRepo:    repository.GetEmployeeRepo{},
	}

	e.GET(utility.APIGet, c.GetEmployees)
	e.POST(utility.APICreate, c.CreateEmployee)
	e.PUT(utility.APIUpdate, c.UpdateEmployee)
	e.DELETE(utility.APIDelete, c.DeleteEmployee)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang")
	})

	e.Logger.Fatal(e.Start(":8090"))
}
