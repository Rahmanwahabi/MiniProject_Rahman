package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"not null;unique"`
	Password string `json:"password" form:"password" gorm:"not null"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Guru', 'Siswa');default:'Guru'"`
	Token    string `json:"token" form:"token"`
}
