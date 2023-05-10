package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func CreateGuru(guru *models.Guru) error {
	if err := config.DB.Save(guru).Error; err != nil {
		return err
	}
	return nil
}

func GetAllGuru() (gurus []models.Guru, err error) {
	if err = config.DB.Find(&gurus).Error; err != nil {
		return
	}
	return
}

func GetGuru(id uint) (guru models.Guru, err error) {
	guru.ID = id
	if err = config.DB.First(&guru).Error; err != nil {
		return
	}
	return
}

func UpdateGuru(guru *models.Guru) error {
	if err := config.DB.Updates(guru).Error; err != nil {
		return err
	}
	return nil
}

func DeleteGuru(guru *models.Guru) error {
	if err := config.DB.Delete(guru).Error; err != nil {
		return err
	}
	return nil
}
