package models

import (
	"gorm.io/gorm"
)

type Mapel struct {
	gorm.Model
	Name      string   `json:"name" form:"name"`
	Kd_Mapel  string   `json:"kd_mapel" form:"kd_mapel" gorm:"unique;not null"`
	GuruID    uint     `json:"guru_id" form:"guru_id"   gorm:"not null"`
	Kelas     tagKType `json:"kelas" form:"kelas" gorm:"type:enum('Ipa', 'Ips');default:'Ipa'"`
	Isi_Mapel string   `json:"isi_mapel" form:"isi_mapel"`
}
