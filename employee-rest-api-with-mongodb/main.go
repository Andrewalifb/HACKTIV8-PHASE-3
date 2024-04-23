package main

import (
	"log"

	"github.com/Andrewalifb/HACKTIV8-PHASE-3/employee-rest-api-mongodb/controller"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.POST("/employee", controller.CreateEmployee)
	e.GET("/all-employee", controller.GetAllEmployee)
	e.GET("/employee/:id", controller.GetEmployeeById)
	e.GET("/paging", controller.GetPagging)
	e.GET("/sorting", controller.GetSorting)
	e.PUT("/employee/:id", controller.UpdateDataEmployee)
	e.DELETE("/employee/:id", controller.DeleteDataPerson)
	e.Logger.Fatal(e.Start(":8080"))
}
