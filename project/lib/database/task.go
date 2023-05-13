package database

import (
	"MiniProject_Rahman/project/config"
	"MiniProject_Rahman/project/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTasks() (tasks []models.Task, err error)
	GetTask(id uint) (task models.Task, err error)
	GetTasksByBoardId(boardID uint) (tasks []models.Task, err error)
	GetTasksByUserId(userID uint) (tasks []models.Task, err error)
	UpdateTask(task *models.Task) error
	DeleteTask(task *models.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) CreateTask(task *models.Task) error {
	if err := config.DB.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetTasks() (tasks []models.Task, err error) {
	if err = config.DB.Find(&tasks).Error; err != nil {
		return
	}
	return
}

func (t *taskRepository) GetTask(id uint) (task models.Task, err error) {
	task.ID = id
	if err = config.DB.First(&task).Error; err != nil {
		return
	}
	return
}

func (t *taskRepository) GetTasksByBoardId(boardID uint) (tasks []models.Task, err error) {
	var task models.Task
	if err := t.db.Where("board_id = ?", boardID).Find(&task).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskRepository) GetTasksByUserId(userID uint) (tasks []models.Task, err error) {
	var task models.Task
	if err := t.db.Where("user_id = ?", userID).Find(&task).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskRepository) UpdateTask(task *models.Task) error {
	if err := config.DB.Updates(task).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) DeleteTask(task *models.Task) error {
	if err := config.DB.Delete(task).Error; err != nil {
		return err
	}
	return nil
}
