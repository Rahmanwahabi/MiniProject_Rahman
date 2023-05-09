package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func CreateSiswa(siswa *models.Siswa) error {
	if err := config.DB.Create(siswa).Error; err != nil {
		return err
	}
	return nil
}

func GetAllSiswa() (siswa []models.Siswa, err error) {
	if err = config.DB.Find(&siswa).Error; err != nil {
		return
	}
	return
}

func GetSiswa(siswa *models.Siswa) (err error) {
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

func LoginSiswa(siswa *models.Siswa) error {
	if err := config.DB.Where("email = ? AND password = ?", siswa.Email, siswa.Password).First(&siswa).Error; err != nil {
		return err
	}
	return nil
}
