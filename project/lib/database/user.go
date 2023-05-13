package database

import (
	"MiniProject_Rahman/project/config"
	"MiniProject_Rahman/project/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUsers() (users []models.User, err error)
	GetUser(user *models.User) (err error)
	GetUserWithBoard(id uint) (user models.User, err error)
	GetUserWithTask(id uint) (user models.User, err error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
	LoginUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUsers() (users []models.User, err error) {
	if err = config.DB.Model(&models.User{}).Preload("Boards").Find(&users).Error; err != nil {
		return
	}
	if err = config.DB.Model(&models.User{}).Preload("Tasks").Find(&users).Error; err != nil {
		return
	}
	return
}

func (u *userRepository) GetUser(user *models.User) (err error) {
	if err = config.DB.First(&user).Error; err != nil {
		return
	}
	return
}

func (u *userRepository) GetUserWithBoard(id uint) (user models.User, err error) {
	user.ID = id
	if err = config.DB.Model(&models.User{}).Preload("Boards").First(&user).Error; err != nil {
		return
	}
	return
}

func (u *userRepository) GetUserWithTask(id uint) (user models.User, err error) {
	user.ID = id
	if err = config.DB.Model(&models.User{}).Preload("Tasks").First(&user).Error; err != nil {
		return
	}
	return
}

func (u *userRepository) UpdateUser(user *models.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) LoginUser(user *models.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}
