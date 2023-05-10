package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func CreateUser(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers() (users []models.User, err error) {
	if err = config.DB.Find(&users).Error; err != nil {
		return
	}
	return
}

func GetUser(user *models.User) (err error) {
	if err = config.DB.First(&user).Error; err != nil {
		return
	}
	return
}

func UpdateUser(user *models.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(user *models.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}
