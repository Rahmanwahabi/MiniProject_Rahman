package routes

import (
	"MiniProject_Rahman/controllers"
	"MiniProject_Rahman/middlewares"
	"MiniProject_Rahman/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New(e *echo.Echo, db *gorm.DB, guruController *controllers.GuruController) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.Use(middlewares.JWT())
	user.GET("", controllers.GetAllUserControllers)
	user.GET("/:id", controllers.GetUserController)
	user.POST("", controllers.CreateUserController)
	user.PUT("/:id", controllers.UpdateUserController)
	user.DELETE("/:id", controllers.DeleteUserController)

	// siswa collection
	// Membuat instance SiswaUseCase
	siswaUseCase := usecase.NewSiswaUseCase()
	// Membuat instance SiswaController dan menyediakan SiswaUseCase
	siswaController := controllers.NewSiswaController(*siswaUseCase)

	// Define route handlers
	siswa := e.Group("/siswa")
	siswa.Use(middlewares.JWT())
	siswa.GET("", siswaController.GetAllSiswa)        // Changed function signature
	siswa.GET("/:id", siswaController.GetSiswa)       // Changed function signature
	siswa.POST("", siswaController.CreateSiswa)       // Changed function signature
	siswa.PUT("/:id", siswaController.UpdateSiswa)    // Changed function signature
	siswa.DELETE("/:id", siswaController.DeleteSiswa) // Changed function signature

	guru := e.Group("/guru")
	guru.Use(middlewares.JWT())
	guru.GET("", guruController.GetAllGuru)
	guru.GET(":id", guruController.GetGuru)
	guru.POST("", guruController.CreateGuru)
	guru.PUT(":id", guruController.UpdateGuru)
	guru.DELETE(":id", guruController.DeleteGuru)

}
