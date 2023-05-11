package usecase

import (
	"MiniProject_Rahman/lib/database"
	"MiniProject_Rahman/models"
)

type SoalUseCase struct{}

func NewSoalUseCase() *SoalUseCase {
	return &SoalUseCase{}
}

func (uc *SoalUseCase) CreateSoal(soal *models.Soal) error {
	err := database.CreateSoal(soal)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *SoalUseCase) GetSoalByID(id uint) (*models.Soal, error) {
	soal, err := database.GetSoalByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}
	return soal, nil
}

func (uc *SoalUseCase) UpdateSoal(soal *models.Soal) error {
	err := database.UpdateSoal(soal)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *SoalUseCase) DeleteSoal(soal *models.Soal) error {
	err := database.DeleteSoal(soal)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *SoalUseCase) GetSoals() ([]models.Soal, error) {
	soals, err := database.GetSoals()
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return soals, err
}
