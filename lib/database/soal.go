package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func CreateSoal(soal *models.Soal) error {
	db := config.DB
	result := db.Create(soal)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetSoals() (soals []models.Soal, err error) {
	if err = config.DB.Find(&soals).Error; err != nil {
		return
	}
	return
}

func GetSoalByID(id uint) (*models.Soal, error) {
	db := config.DB
	soal := new(models.Soal)
	result := db.First(soal, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return soal, nil
}

func UpdateSoal(soal *models.Soal) error {
	db := config.DB
	result := db.Save(soal)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteSoal(soal *models.Soal) error {
	db := config.DB
	result := db.Delete(soal)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
