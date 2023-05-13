package usecase

import (
	"MiniProject_Rahman/project/lib/database"
	"MiniProject_Rahman/project/models"
	"errors"
	"fmt"
)

type TaskUsecase interface {
	CreateTask(task *models.Task) error
	GetTask(id uint) (task models.Task, err error)
	GetListTasks() (tasks []models.Task, err error)
	UpdateTask(task *models.Task) (err error)
	DeleteTask(id uint) (err error)
}

type taskUsecase struct {
	taskRepository database.TaskRepository
}

func NewTaskUsecase(taskRepo database.TaskRepository) *taskUsecase {
	return &taskUsecase{taskRepository: taskRepo}
}

func (t *taskUsecase) CreateTask(task *models.Task) error {
	// check name cannot be empty
	if task.TaskName == "" {
		return errors.New("task title cannot be empty")
	}

	// check description
	if task.TaskDescription == "" {
		return errors.New("task description cannot be empty")
	}

	err := t.taskRepository.CreateTask(task)
	if err != nil {
		return err
	}

	return nil

}

func (t *taskUsecase) GetTask(id uint) (task models.Task, err error) {
	task, err = t.taskRepository.GetTask(id)
	if err != nil {
		fmt.Println("GetTask: Error getting task from database")
		return
	}
	return
}

func (t *taskUsecase) GetListTasks() (tasks []models.Task, err error) {
	tasks, err = t.taskRepository.GetTasks()
	if err != nil {
		fmt.Println("GetListTasks: Error getting tasks from database")
		return
	}
	return
}

func (t *taskUsecase) UpdateTask(task *models.Task) (err error) {
	err = t.taskRepository.UpdateTask(task)
	if err != nil {
		fmt.Println("UpdateTask : Error updating task, err: ", err)
		return
	}
	return

}

func (t *taskUsecase) DeleteTask(id uint) (err error) {
	task := models.Task{}
	task.ID = id
	err = t.taskRepository.DeleteTask(&task)
	if err != nil {
		fmt.Println("DeleteTask : error deleting task, err: ", err)
		return
	}
	return
}
