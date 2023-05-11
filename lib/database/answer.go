package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func CreateAnswer(answer *models.Answer) error {
	db := config.DB
	result := db.Create(answer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAnswers() (answers []models.Answer, err error) {
	if err = config.DB.Find(&answers).Error; err != nil {
		return
	}
	return
}

func GetAnswerByID(id uint) (*models.Answer, error) {
	db := config.DB
	answer := new(models.Answer)
	result := db.First(answer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return answer, nil
}

func UpdateAnswer(answer *models.Answer) error {
	db := config.DB
	result := db.Save(answer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteAnswer(answer *models.Answer) error {
	db := config.DB
	result := db.Delete(answer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
