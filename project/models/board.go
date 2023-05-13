package models

import (
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	BoardName        string `json:"board_name" form:"board_name"`
	BoardDescription string `json:"board_description" form:"board_description"`
	UserID           uint   `json:"user_id" form:"user_id" gorm:"index;not null"`
	Tasks            []Task
}
