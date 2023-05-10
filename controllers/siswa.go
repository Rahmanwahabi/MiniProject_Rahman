package controllers

import (
	"MiniProject_Rahman/models"
	"MiniProject_Rahman/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// SiswaController adalah struktur controller untuk entitas Siswa
type SiswaController struct {
	SiswaUseCase usecase.SiswaUseCase
}

// NewSiswaController adalah fungsi untuk membuat instance baru dari SiswaController
func NewSiswaController(siswaUseCase usecase.SiswaUseCase) *SiswaController {
	return &SiswaController{siswaUseCase}
}

// GetAllSiswa adalah handler untuk mengambil semua data siswa
func (controller *SiswaController) GetAllSiswa(c echo.Context) error {
	siswas, err := controller.SiswaUseCase.GetAllSiswa()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, siswas)
}

// GetSiswa adalah handler untuk mengambil data siswa berdasarkan ID
func (controller *SiswaController) GetSiswa(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	siswa, err := controller.SiswaUseCase.GetSiswa(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, siswa)
}

// CreateSiswa adalah handler untuk membuat data siswa baru
func (controller *SiswaController) CreateSiswa(c echo.Context) error {
	var siswa models.Siswa
	err := c.Bind(&siswa)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err = controller.SiswaUseCase.CreateSiswa(&siswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, siswa)
}

// UpdateSiswa adalah handler untuk mengupdate data siswa
func (controller *SiswaController) UpdateSiswa(c echo.Context) error {
	var siswa models.Siswa
	err := c.Bind(&siswa)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err = controller.SiswaUseCase.UpdateSiswa(&siswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, siswa)
}

// DeleteSiswa adalah handler untuk menghapus data siswa
func (controller *SiswaController) DeleteSiswa(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	siswa, err := controller.SiswaUseCase.GetSiswa(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = controller.SiswaUseCase.DeleteSiswa(&siswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
