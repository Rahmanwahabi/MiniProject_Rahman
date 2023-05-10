package usecase

import (
	"MiniProject_Rahman/lib/database"
	"MiniProject_Rahman/models"
	"errors"
)

type SiswaUseCase struct{}

func NewSiswaUseCase() *SiswaUseCase {
	return &SiswaUseCase{}
}

func (uc *SiswaUseCase) GetAllSiswa() ([]models.Siswa, error) {
	return database.GetAllSiswa()
}

func (uc *SiswaUseCase) GetSiswa(id uint) (models.Siswa, error) {
	if id == 0 {
		return models.Siswa{}, errors.New("invalid ID")
	}
	return database.GetSiswa(id)
}

func (uc *SiswaUseCase) UpdateSiswa(siswa *models.Siswa) error {
	if err := validateSiswa(siswa); err != nil {
		return err
	}
	return database.UpdateSiswa(siswa)
}

func (uc *SiswaUseCase) DeleteSiswa(siswa *models.Siswa) error {
	if err := validateSiswa(siswa); err != nil {
		return err
	}
	return database.DeleteSiswa(siswa)
}

func (uc *SiswaUseCase) CreateSiswa(siswa *models.Siswa) error {
	if err := validateSiswa(siswa); err != nil {
		return err
	}
	return database.CreateSiswa(siswa)
}

func validateSiswa(siswa *models.Siswa) error {
	if siswa == nil {
		return errors.New("siswa is nil")
	}
	if siswa.Name == "" {
		return errors.New("siswa name is required")
	}
	if siswa.Email == "" {
		return errors.New("siswa email is required")
	}
	return nil
}
