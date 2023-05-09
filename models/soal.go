package models

import (
	"gorm.io/gorm"
)

type Soal struct {
	gorm.Model
	Name     string   `json:"name" form:"name"`
	Kd_Soal  string   `json:"kd_soal" form:"kd_soal" gorm:"unique;not null"`
	GuruID   uint     `json:"guru_id" form:"guru_id"   gorm:"not null"`
	Kelas    tagKType `json:"kelas" form:"kelas" gorm:"type:enum('Ipa', 'Ips');default:'Ipa'"`
	Isi_soal string   `json:"isi_soal" form:"isi_soal"`
}
