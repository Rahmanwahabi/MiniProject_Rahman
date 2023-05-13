package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName        string `json:"task_name" form:"task_name"`
	TaskDescription string `json:"task_description" form:"task_description"`
	DueDate         string `json:"due_date" from:"due_date"`
	Complete        bool   `json:"complate" from:"complate" gorm:"not null"`
	CompleteDate    string `gorm:"not null"`
	UserID          uint   `json:"user_id" form:"user_id" gorm:"index"`
	BoardID         uint   `json:"board_id" form:"board_id" gorm:"index"`
}
