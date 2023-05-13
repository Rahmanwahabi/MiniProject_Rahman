package main

import (
	"MiniProject_Rahman/project/config"
	"MiniProject_Rahman/project/middlewares"
	"MiniProject_Rahman/project/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	e.Use(middlewares.Logger())

	routes.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
