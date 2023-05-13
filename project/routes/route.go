package routes

import (
	"MiniProject_Rahman/project/constant"
	"MiniProject_Rahman/project/controllers"
	"MiniProject_Rahman/project/lib/database"
	"MiniProject_Rahman/project/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func NewRoute(e *echo.Echo, db *gorm.DB) {

	userRepository := database.NewUserRepository(db)
	boardRepository := database.NewBoardRepository(db)
	taskRepository := database.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, boardRepository, taskRepository)
	boardUsecase := usecase.NewBoardUsecase(boardRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)

	authController := controllers.NewAuthController(userUsecase)
	userController := controllers.NewUserController(userUsecase, userRepository)
	boardController := controllers.NewBoardController(boardUsecase, boardRepository)
	taskController := controllers.NewTaskController(taskUsecase, taskRepository)

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", authController.LoginUserController)
	e.POST("/register", userController.CreateUserController)

	// user collection
	user := e.Group("/users", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.GET("", userController.GetUserscontroller)
	user.GET("/:id", userController.GetUserController)
	user.POST("", userController.CreateUserController)
	user.PUT("/:id", userController.UpdateUserController)
	user.DELETE("/:id", userController.DeleteUserController)

	// board collection
	board := e.Group("/boards", middleware.JWT([]byte(constant.SECRET_JWT)))
	board.GET("", boardController.GetBoardsController)
	board.GET("/:id", boardController.GetBoardController)
	board.POST("", boardController.CreateBoardController)
	board.PUT("/:id", boardController.UpdateBoardController)
	board.DELETE("/:id", boardController.DeleteBoardController)

	// task collection
	task := e.Group("/tasks", middleware.JWT([]byte(constant.SECRET_JWT)))
	task.GET("", taskController.GetTasksController)
	task.GET("/:id", taskController.GetTaskController)
	task.POST("", taskController.CreateTaskController)
	task.PUT("/:id", taskController.UpdateTaskController)
	task.DELETE("/:id", taskController.DeleteTaskController)

}
