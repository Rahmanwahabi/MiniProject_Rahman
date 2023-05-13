package controllers

import (
	"MiniProject_Rahman/project/lib/database"
	"MiniProject_Rahman/project/models"
	"MiniProject_Rahman/project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskController interface {
	GetTasksController(c echo.Context) error
	GetTaskController(c echo.Context) error
	CreateTaskController(c echo.Context) error
	DeleteTaskController(c echo.Context) error
	UpdateTaskController(c echo.Context) error
}

type taskController struct {
	taskUsecase    usecase.TaskUsecase
	taskRepository database.TaskRepository
}

func NewTaskController(
	taskUsecase usecase.TaskUsecase,
	taskRepository database.TaskRepository,
) *taskController {
	return &taskController{
		taskUsecase,
		taskRepository,
	}
}

func (t *taskController) GetTasksController(c echo.Context) error {
	tasks, e := t.taskRepository.GetTasks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"tasks":  tasks,
	})
}

func (t *taskController) GetTaskController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	task, err := t.taskRepository.GetTask(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get task",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"tasks":  task,
	})
}

// create new task
func (t *taskController) CreateTaskController(c echo.Context) error {
	task := models.Task{}
	c.Bind(&task)

	if err := t.taskUsecase.CreateTask(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create task",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new task",
		"task":    task,
	})
}

// delete task by id
func (t *taskController) DeleteTaskController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := t.taskUsecase.DeleteTask(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete task",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf task tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete task",
	})
}

// update task by id
func (t *taskController) UpdateTaskController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	task := models.Task{}
	c.Bind(&task)
	task.ID = uint(id)
	if err := t.taskUsecase.UpdateTask(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update task",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf task tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update task",
	})
}
