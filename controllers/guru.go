package controllers

import (
	"MiniProject_Rahman/models"
	"MiniProject_Rahman/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GuruController struct {
	useCase *usecase.GuruUseCase
}

func NewGuruController(useCase *usecase.GuruUseCase) *GuruController {
	return &GuruController{
		useCase: useCase,
	}
}

func (c *GuruController) CreateGuru(ctx echo.Context) error {
	guru := new(models.Guru)
	if err := ctx.Bind(guru); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err := c.useCase.CreateGuru(guru)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create guru")
	}

	return ctx.JSON(http.StatusCreated, guru)
}

func (c *GuruController) GetAllGuru(ctx echo.Context) error {
	gurus, err := c.useCase.GetAllGuru()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to fetch gurus")
	}

	return ctx.JSON(http.StatusOK, gurus)
}

func (c *GuruController) GetGuru(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid guru ID")
	}

	guru, err := c.useCase.GetGuru(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Guru not found")
	}

	return ctx.JSON(http.StatusOK, guru)
}

func (c *GuruController) UpdateGuru(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid guru ID")
	}

	guru := new(models.Guru)
	if err := ctx.Bind(guru); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	guru.ID = uint(id)

	err = c.useCase.UpdateGuru(guru)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update guru")
	}

	return ctx.JSON(http.StatusOK, guru)
}

func (controller *GuruController) DeleteGuru(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	guru, err := controller.useCase.GetGuru(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = controller.useCase.DeleteGuru(&guru)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
