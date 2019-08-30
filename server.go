package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"GO-ECHO-TUTORIAL/utility"
	"GO-ECHO-TUTORIAL/controller"
	"GO-ECHO-TUTORIAL/service"
	"GO-ECHO-TUTORIAL/repository"
)

func main() {
	serve()
}

func serve() {
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

	initRoutes(e, c)

	e.Logger.Fatal(e.Start(":8090"))
}

func initRoutes(e *echo.Echo, c controller.EmpController) {
	g := e.Group("/go-echo-yudi")
	g.GET(utility.APIGet, c.GetEmployees)
	g.POST(utility.APICreate, c.CreateEmployee)
	g.PUT(utility.APIUpdate, c.UpdateEmployee)
	g.DELETE(utility.APIDelete, c.DeleteEmployee)
}