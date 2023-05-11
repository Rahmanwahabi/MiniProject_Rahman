package controllers

import (
	"MiniProject_Rahman/models"
	"MiniProject_Rahman/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MapelController struct {
	useCase *usecase.MapelUseCase
}

func NewMapelController(useCase *usecase.MapelUseCase) *MapelController {
	return &MapelController{
		useCase: useCase,
	}
}

func (c *MapelController) CreateMapel(ctx echo.Context) error {
	mapel := new(models.Mapel)
	if err := ctx.Bind(mapel); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err := c.useCase.CreateMapel(mapel)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create mapel")
	}

	return ctx.JSON(http.StatusCreated, mapel)
}

func (c *MapelController) GetMapels(ctx echo.Context) error {
	mapels, err := c.useCase.GetMapels()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to fetch mapels")
	}

	return ctx.JSON(http.StatusOK, mapels)
}

func (c *MapelController) GetMapel(ctx echo.Context) error {
	id := ctx.Param("id")

	mapelID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid mapel ID")
	}

	mapel, err := c.useCase.GetMapelByID(uint(mapelID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Mapel not found")
	}

	return ctx.JSON(http.StatusOK, mapel)
}

func (c *MapelController) UpdateMapel(ctx echo.Context) error {
	id := ctx.Param("id")

	mapelID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid mapel ID")
	}

	mapel := new(models.Mapel)
	if err := ctx.Bind(mapel); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	mapel.ID = uint(mapelID)

	err = c.useCase.UpdateMapel(mapel)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update mapel")
	}

	return ctx.JSON(http.StatusOK, mapel)
}

func (c *MapelController) DeleteMapel(ctx echo.Context) error {
	id := ctx.Param("id")

	mapelID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid mapel ID")
	}

	mapel := new(models.Mapel)
	mapel.ID = uint(mapelID)

	err = c.useCase.DeleteMapel(mapel)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete mapel")
	}

	return ctx.NoContent(http.StatusNoContent)
}
