package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type tagKType string

const (
	Ipa tagKType = "Ipa"
	Ips tagKType = "Ips"
)

func (t *tagKType) Scan(value interface{}) error {
	*t = tagKType(value.([]byte))
	return nil
}

func (t tagKType) Value() (driver.Value, error) {
	return string(t), nil
}

type Kelas struct {
	gorm.Model
	Nm_Kelas tagKType `json:"nm_kelas" form:"nm_kelas" gorm:"type:enum('Ipa', 'Ips');default:'Ipa'"`
	Kd_Kelas string   `json:"kd_kelas" form:"kd_kelas"`
}
