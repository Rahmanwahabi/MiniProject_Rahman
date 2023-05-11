package controllers

import (
	"MiniProject_Rahman/models"
	"MiniProject_Rahman/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KelasController struct {
	useCase *usecase.KelasUseCase
}

func NewKelasController(useCase *usecase.KelasUseCase) *KelasController {
	return &KelasController{
		useCase: useCase,
	}
}

func (c *KelasController) CreateKelas(ctx echo.Context) error {
	kelas := new(models.Kelas)
	if err := ctx.Bind(kelas); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err := c.useCase.CreateKelas(kelas)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create kelas")
	}

	return ctx.JSON(http.StatusCreated, kelas)
}

func (c *KelasController) GetKelases(ctx echo.Context) error {
	kelases, err := c.useCase.GetKelases()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to fetch kelases")
	}

	return ctx.JSON(http.StatusOK, kelases)
}

func (c *KelasController) GetKelas(ctx echo.Context) error {
	id := ctx.Param("id")

	kelasID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid kelas ID")
	}

	kelas, err := c.useCase.GetKelasByID(uint(kelasID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Kelas not found")
	}

	return ctx.JSON(http.StatusOK, kelas)
}

func (c *KelasController) UpdateKelas(ctx echo.Context) error {
	id := ctx.Param("id")

	kelasID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid kelas ID")
	}

	kelas := new(models.Kelas)
	if err := ctx.Bind(kelas); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	kelas.ID = uint(kelasID)

	err = c.useCase.UpdateKelas(kelas)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update kelas")
	}

	return ctx.JSON(http.StatusOK, kelas)
}

func (c *KelasController) DeleteKelas(ctx echo.Context) error {
	id := ctx.Param("id")

	kelasID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid kelas ID")
	}

	kelas := new(models.Kelas)
	kelas.ID = uint(kelasID)

	err = c.useCase.DeleteKelas(kelas)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete kelas")
	}

	return ctx.NoContent(http.StatusNoContent)
}
