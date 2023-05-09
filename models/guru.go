package models

import (
	"gorm.io/gorm"
)

type Guru struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"not null;unique"`
	Password string `json:"password" form:"password" gorm:"not null"`
	Nip      string `json:"nip" form:"nip"`
	Jk       string `json:"jk" form:"jk" gorm:"type:enum('Laki-laki', 'Perempuan');default:'Laki-laki'"`
	Token    string `json:"token" form:"token"`
}
