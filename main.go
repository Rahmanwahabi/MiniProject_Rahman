package main

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/controllers"
	"MiniProject_Rahman/middlewares"
	"MiniProject_Rahman/routes"
	"MiniProject_Rahman/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	e.Use(middlewares.Logger())

	// Inisialisasi use case
	useCase := usecase.NewGuruUseCase() // Inisialisasi use case
	// Inisialisasi GuruController
	guruController := controllers.NewGuruController(useCase)

	routes.New(e, db, guruController)

	e.Logger.Fatal(e.Start(":8080"))
}
