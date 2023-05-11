package usecase

import (
	"MiniProject_Rahman/lib/database"
	"MiniProject_Rahman/models"
)

type KelasUseCase struct{}

func NewKelasUseCase() *KelasUseCase {
	return &KelasUseCase{}
}

func (uc *KelasUseCase) CreateKelas(kelas *models.Kelas) error {
	err := database.CreateKelas(kelas)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *KelasUseCase) GetKelasByID(id uint) (*models.Kelas, error) {
	kelas, err := database.GetKelasByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}
	return kelas, nil
}

func (uc *KelasUseCase) UpdateKelas(kelas *models.Kelas) error {
	err := database.UpdateKelas(kelas)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *KelasUseCase) DeleteKelas(kelas *models.Kelas) error {
	err := database.DeleteKelas(kelas)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *KelasUseCase) GetKelases() ([]models.Kelas, error) {
	kelases, err := database.GetKelases()
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return kelases, err
}
