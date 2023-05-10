package usecase

import (
	"MiniProject_Rahman/lib/database"
	"MiniProject_Rahman/models"
)

type GuruUseCase struct {
}

func NewGuruUseCase() *GuruUseCase {
	return &GuruUseCase{}
}

func (uc *GuruUseCase) CreateGuru(guru *models.Guru) error {
	err := database.CreateGuru(guru)
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return err
}

func (uc *GuruUseCase) GetAllGuru() ([]models.Guru, error) {
	gurus, err := database.GetAllGuru()
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return gurus, err
}

func (uc *GuruUseCase) GetGuru(id uint) (models.Guru, error) {
	guru, err := database.GetGuru(id)
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return guru, err
}

func (uc *GuruUseCase) UpdateGuru(guru *models.Guru) error {
	err := database.UpdateGuru(guru)
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return err
}

func (uc *GuruUseCase) DeleteGuru(guru *models.Guru) error {
	err := database.DeleteGuru(guru)
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return err
}
