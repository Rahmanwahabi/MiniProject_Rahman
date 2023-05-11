package usecase

import (
	"MiniProject_Rahman/lib/database"
	"MiniProject_Rahman/models"
)

type RoleUseCase struct{}

func NewRoleUseCase() *RoleUseCase {
	return &RoleUseCase{}
}

func (uc *RoleUseCase) CreateRole(role *models.Role) error {
	err := database.CreateRole(role)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *RoleUseCase) GetRoleByID(id uint) (*models.Role, error) {
	role, err := database.GetRoleByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}
	return &role, nil
}

func (uc *RoleUseCase) UpdateRole(role *models.Role) error {
	err := database.UpdateRole(role)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *RoleUseCase) DeleteRole(role *models.Role) error {
	err := database.DeleteRole(role)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *RoleUseCase) GetRoles() ([]models.Role, error) {
	roles, err := database.GetRoles()
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return roles, err
}
