package usecase

import (
	"MiniProject_Rahman/lib/database"
	"MiniProject_Rahman/models"
)

type MapelUseCase struct{}

func NewMapelUseCase() *MapelUseCase {
	return &MapelUseCase{}
}

func (uc *MapelUseCase) CreateMapel(mapel *models.Mapel) error {
	err := database.CreateMapel(mapel)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *MapelUseCase) GetMapelByID(id uint) (*models.Mapel, error) {
	mapel, err := database.GetMapelByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}
	return mapel, nil
}

func (uc *MapelUseCase) UpdateMapel(mapel *models.Mapel) error {
	err := database.UpdateMapel(mapel)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *MapelUseCase) DeleteMapel(mapel *models.Mapel) error {
	err := database.DeleteMapel(mapel)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *MapelUseCase) GetMapels() ([]models.Mapel, error) {
	mapels, err := database.GetMapels()
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return mapels, err
}
