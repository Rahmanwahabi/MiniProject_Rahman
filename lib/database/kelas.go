package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func CreateKelas(kelas *models.Kelas) error {
	db := config.DB
	result := db.Create(kelas)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetKelases() (kelases []models.Kelas, err error) {
	if err = config.DB.Find(&kelases).Error; err != nil {
		return
	}
	return
}

func GetKelasByID(id uint) (*models.Kelas, error) {
	db := config.DB
	kelas := new(models.Kelas)
	result := db.First(kelas, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return kelas, nil
}

func UpdateKelas(kelas *models.Kelas) error {
	db := config.DB
	result := db.Save(kelas)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteKelas(kelas *models.Kelas) error {
	db := config.DB
	result := db.Delete(kelas)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
