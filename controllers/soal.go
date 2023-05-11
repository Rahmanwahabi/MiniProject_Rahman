package controllers

import (
	"MiniProject_Rahman/models"
	"MiniProject_Rahman/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SoalController struct {
	useCase *usecase.SoalUseCase
}

func NewSoalController(useCase *usecase.SoalUseCase) *SoalController {
	return &SoalController{
		useCase: useCase,
	}
}

func (c *SoalController) CreateSoal(ctx echo.Context) error {
	soal := new(models.Soal)
	if err := ctx.Bind(soal); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err := c.useCase.CreateSoal(soal)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create soal")
	}

	return ctx.JSON(http.StatusCreated, soal)
}

func (c *SoalController) GetSoals(ctx echo.Context) error {
	soals, err := c.useCase.GetSoals()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to fetch soals")
	}

	return ctx.JSON(http.StatusOK, soals)
}

func (c *SoalController) GetSoal(ctx echo.Context) error {
	id := ctx.Param("id")

	soalID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid soal ID")
	}

	soal, err := c.useCase.GetSoalByID(uint(soalID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Soal not found")
	}

	return ctx.JSON(http.StatusOK, soal)
}

func (c *SoalController) UpdateSoal(ctx echo.Context) error {
	id := ctx.Param("id")

	soalID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid soal ID")
	}

	soal := new(models.Soal)
	if err := ctx.Bind(soal); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	soal.ID = uint(soalID)

	err = c.useCase.UpdateSoal(soal)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update soal")
	}

	return ctx.JSON(http.StatusOK, soal)
}

func (c *SoalController) DeleteSoal(ctx echo.Context) error {
	id := ctx.Param("id")

	soalID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid soal ID")
	}

	soal := new(models.Soal)
	soal.ID = uint(soalID)

	err = c.useCase.DeleteSoal(soal)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete soal")
	}

	return ctx.NoContent(http.StatusNoContent)
}
