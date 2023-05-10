package models

import (
	"gorm.io/gorm"
)

type Siswa struct {
	gorm.Model
	Name     string   `json:"name" form:"name"`
	Email    string   `json:"email" form:"email" gorm:"not null;unique"`
	Password string   `json:"password" form:"password" gorm:"not null"`
	Nis      string   `json:"nis" form:"nis"`
	Kelas    tagKType `json:"kelas" form:"kelas" gorm:"type:enum('Ipa', 'Ips');default:'Ipa'"`
	Jk       string   `json:"jk" form:"jk" gorm:"type:enum('Laki-laki', 'Perempuan');default:'Laki-laki'"`
}
