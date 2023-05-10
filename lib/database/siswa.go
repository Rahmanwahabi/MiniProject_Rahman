package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func GetAllSiswa() (siswas []models.Siswa, err error) {
	if err = config.DB.Find(&siswas).Error; err != nil {
		return
	}
	return
}

func GetSiswa(id uint) (siswa models.Siswa, err error) {
	siswa.ID = id
	if err = config.DB.First(&siswa).Error; err != nil {
		return
	}
	return
}

func UpdateSiswa(siswa *models.Siswa) error {
	if err := config.DB.Updates(siswa).Error; err != nil {
		return err

	}
	return nil
}
func DeleteSiswa(siswa *models.Siswa) error {
	if err := config.DB.Delete(siswa).Error; err != nil {
		return err

	}
	return nil
}

func CreateSiswa(siswa *models.Siswa) error {
	if arr := config.DB.Create(siswa).Error; arr != nil {
		return arr

	}
	return nil
}
