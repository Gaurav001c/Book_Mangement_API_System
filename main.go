package main

import (
	"github.com/GAURAV/BookApiTask/database"
	"github.com/GAURAV/BookApiTask/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	database.ConnectDB()

	e := echo.New()
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":9000"))

}
