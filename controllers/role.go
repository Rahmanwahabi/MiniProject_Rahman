package controllers

import (
	"MiniProject_Rahman/models"
	"MiniProject_Rahman/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleController struct {
	useCase *usecase.RoleUseCase
}

func NewRoleController(useCase *usecase.RoleUseCase) *RoleController {
	return &RoleController{
		useCase: useCase,
	}
}

func (c *RoleController) CreateRole(ctx echo.Context) error {
	role := new(models.Role)
	if err := ctx.Bind(role); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err := c.useCase.CreateRole(role)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create role")
	}

	return ctx.JSON(http.StatusCreated, role)
}

func (c *RoleController) GetRoles(ctx echo.Context) error {
	roles, err := c.useCase.GetRoles()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to fetch roles")
	}

	return ctx.JSON(http.StatusOK, roles)
}

func (c *RoleController) GetRole(ctx echo.Context) error {
	id := ctx.Param("id")

	roleID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid role ID")
	}

	role, err := c.useCase.GetRoleByID(uint(roleID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Role not found")
	}

	return ctx.JSON(http.StatusOK, role)
}

func (c *RoleController) UpdateRole(ctx echo.Context) error {
	id := ctx.Param("id")

	roleID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid role ID")
	}

	role := new(models.Role)
	if err := ctx.Bind(role); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	role.ID = uint(roleID)

	err = c.useCase.UpdateRole(role)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update role")
	}

	return ctx.JSON(http.StatusOK, role)
}

func (c *RoleController) DeleteRole(ctx echo.Context) error {
	id := ctx.Param("id")

	roleID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid role ID")
	}

	role := new(models.Role)
	role.ID = uint(roleID)

	err = c.useCase.DeleteRole(role)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete role")
	}

	return ctx.NoContent(http.StatusNoContent)
}
