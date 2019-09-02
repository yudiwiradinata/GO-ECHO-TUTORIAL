package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"

	"GO-ECHO-TUTORIAL/controller"
	"GO-ECHO-TUTORIAL/repository"
	"GO-ECHO-TUTORIAL/service"
	"GO-ECHO-TUTORIAL/utility"
)

type Config struct {
	DataSourceName string
	DriverName     string
}

func NewConfig() *Config {
	return &Config{
		DataSourceName: "root:123456@tcp(localhost:3306)/test",
		DriverName:     "mysql",
	}
}

type Server struct {
	echo               *echo.Echo
	config             *Config
	employeeController *controller.EmployeeController
}

func NewServer(
	config *Config,
	employeeController *controller.EmployeeController,
) *Server {
	s := &Server{
		echo:               echo.New(),
		config:             config,
		employeeController: employeeController,
	}

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	initRoutes(s.echo, s.employeeController)
	return s
}

func ConnectDatabase(config *Config) (*sql.DB, error) {
	return sql.Open(config.DriverName, config.DataSourceName)
}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewConfig)
	container.Provide(ConnectDatabase)
	container.Provide(repository.NewEmployeeRepository)
	container.Provide(service.NewEmployeeService)
	container.Provide(NewServer)
	container.Provide(repository.NewEmployeeRepo)
	container.Provide(service.NewEmployee)
	container.Provide(controller.NewEmployeeController)

	return container
}

func (server *Server) Run() {
	server.echo.Logger.Fatal(server.echo.Start(":8090"))
}

func main() {
	container := BuildContainer()

	err := container.Invoke(func(server *Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}

/*
func serve() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	c := controller.EmpController{
		EmpService: service.EmployeeService{},
		EmpRepo:    repository.EmployeeRepository{},
	}

	initRoutes(e, c)

	e.Logger.Fatal(e.Start(":8090"))

}
*/

func initRoutes(e *echo.Echo, c *controller.EmployeeController) {
	g := e.Group("/go-echo-yudi")
	g.GET(utility.APIGet, c.GetEmployees)
	g.GET(utility.APIGetByName, c.GetEmployeeByName)
	g.POST(utility.APICreate, c.CreateEmployee)
	g.PUT(utility.APIUpdate, c.UpdateEmployee)
	g.DELETE(utility.APIDelete, c.DeleteEmployee)

}
