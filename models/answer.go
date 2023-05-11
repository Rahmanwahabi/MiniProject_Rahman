package models

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	Name       string   `json:"name" form:"name"`
	NameSiswa  string   `json:"name_siswa" form:"name_siswa"`
	Kelas      tagKType `json:"kelas" form:"kelas" gorm:"type:enum('Ipa', 'Ips');default:'Ipa'"`
	Isi_Answer string   `json:"isi_answer" form:"isi_answer"`
}
