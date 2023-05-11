package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func CreateMapel(mapel *models.Mapel) error {
	db := config.DB
	result := db.Create(mapel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetMapels() (mapels []models.Mapel, err error) {
	if err = config.DB.Find(&mapels).Error; err != nil {
		return
	}
	return
}

func GetMapelByID(id uint) (*models.Mapel, error) {
	db := config.DB
	mapel := new(models.Mapel)
	result := db.First(mapel, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return mapel, nil
}

func UpdateMapel(mapel *models.Mapel) error {
	db := config.DB
	result := db.Save(mapel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteMapel(mapel *models.Mapel) error {
	db := config.DB
	result := db.Delete(mapel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
