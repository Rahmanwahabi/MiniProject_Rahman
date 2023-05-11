package database

import (
	"MiniProject_Rahman/config"
	"MiniProject_Rahman/models"
)

func CreateRole(role *models.Role) error {
	if err := config.DB.Save(role).Error; err != nil {
		return err
	}
	return nil
}

func GetRoles() (roles []models.Role, err error) {
	if err = config.DB.Find(&roles).Error; err != nil {
		return
	}
	return
}

func GetRoleByID(id uint) (role models.Role, err error) {
	role.ID = id
	if err = config.DB.First(&role).Error; err != nil {
		return
	}
	return
}

func UpdateRole(role *models.Role) error {
	if err := config.DB.Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRole(role *models.Role) error {
	if err := config.DB.Delete(role).Error; err != nil {
		return err
	}
	return nil
}
