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

	// Membuat instance dari usecase
	siswaUseCase := usecase.NewSiswaUseCase()
	roleUseCase := usecase.NewRoleUseCase()
	kelasUseCase := usecase.NewKelasUseCase()
	mapelUseCase := usecase.NewMapelUseCase()
	soalUseCase := usecase.NewSoalUseCase()
	answerUseCase := usecase.NewAnswerUseCase()

	// Membuat instance dari Controller dengan menginject UseCase
	siswaController := controllers.NewSiswaController(*siswaUseCase)
	roleController := controllers.NewRoleController(roleUseCase)
	kelasController := controllers.NewKelasController(kelasUseCase)
	mapelController := controllers.NewMapelController(mapelUseCase)
	soalController := controllers.NewSoalController(soalUseCase)
	answerController := controllers.NewAnswerController(answerUseCase)

	// user collection
	user := e.Group("/users")
	user.Use(middlewares.JWT())
	user.GET("", controllers.GetAllUserControllers)
	user.GET("/:id", controllers.GetUserController)
	user.POST("", controllers.CreateUserController)
	user.PUT("/:id", controllers.UpdateUserController)
	user.DELETE("/:id", controllers.DeleteUserController)

	// siswa collection
	siswa := e.Group("/siswa")
	siswa.Use(middlewares.JWT())
	siswa.GET("", siswaController.GetAllSiswa)        // Changed function signature
	siswa.GET("/:id", siswaController.GetSiswa)       // Changed function signature
	siswa.POST("", siswaController.CreateSiswa)       // Changed function signature
	siswa.PUT("/:id", siswaController.UpdateSiswa)    // Changed function signature
	siswa.DELETE("/:id", siswaController.DeleteSiswa) // Changed function signature

	// guru collection
	guru := e.Group("/guru")
	guru.Use(middlewares.JWT())
	guru.GET("", guruController.GetAllGuru)
	guru.GET("/:id", guruController.GetGuru)
	guru.POST("", guruController.CreateGuru)
	guru.PUT("/:id", guruController.UpdateGuru)
	guru.DELETE("/:id", guruController.DeleteGuru)

	// role collection

	role := e.Group("/roles")
	role.Use(middlewares.JWT())
	role.GET("", roleController.GetRoles)
	role.GET("/:id", roleController.GetRole)
	role.POST("", roleController.CreateRole)
	role.PUT("/:id", roleController.UpdateRole)
	role.DELETE("/:id", roleController.DeleteRole)

	// kelas collection

	kelas := e.Group("/kelas")
	kelas.Use(middlewares.JWT())
	kelas.GET("", kelasController.GetKelases)
	kelas.GET("/:id", kelasController.GetKelas)
	kelas.POST("", kelasController.CreateKelas)
	kelas.PUT("/:id", kelasController.UpdateKelas)
	kelas.DELETE("/:id", kelasController.DeleteKelas)

	// mapel collection

	mapel := e.Group("/mapel")
	mapel.Use(middlewares.JWT())
	mapel.GET("", mapelController.GetMapels)
	mapel.GET("/:id", mapelController.GetMapel)
	mapel.POST("", mapelController.CreateMapel)
	mapel.PUT("/:id", mapelController.UpdateMapel)
	mapel.DELETE("/:id", mapelController.DeleteMapel)

	// soal collection

	soal := e.Group("/soal")
	soal.Use(middlewares.JWT())
	soal.GET("", soalController.GetSoals)
	soal.GET("/:id", soalController.GetSoal)
	soal.POST("", soalController.CreateSoal)
	soal.PUT("/:id", soalController.UpdateSoal)
	soal.DELETE("/:id", soalController.DeleteSoal)

	// answer collection

	answer := e.Group("/answer")
	answer.Use(middlewares.JWT())
	answer.GET("", answerController.GetAnswers)
	answer.GET("/:id", answerController.GetAnswer)
	answer.POST("", answerController.CreateAnswer)
	answer.PUT("/:id", answerController.UpdateAnswer)
	answer.DELETE("/:id", answerController.DeleteAnswer)

}
