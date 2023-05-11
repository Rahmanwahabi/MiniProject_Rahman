package controllers

import (
	"MiniProject_Rahman/models"
	"MiniProject_Rahman/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AnswerController struct {
	useCase *usecase.AnswerUseCase
}

func NewAnswerController(useCase *usecase.AnswerUseCase) *AnswerController {
	return &AnswerController{
		useCase: useCase,
	}
}

func (c *AnswerController) CreateAnswer(ctx echo.Context) error {
	answer := new(models.Answer)
	if err := ctx.Bind(answer); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err := c.useCase.CreateAnswer(answer)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create answer")
	}

	return ctx.JSON(http.StatusCreated, answer)
}

func (c *AnswerController) GetAnswers(ctx echo.Context) error {
	answers, err := c.useCase.GetAnswers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to fetch answers")
	}

	return ctx.JSON(http.StatusOK, answers)
}

func (c *AnswerController) GetAnswer(ctx echo.Context) error {
	id := ctx.Param("id")

	answerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid answer ID")
	}

	answer, err := c.useCase.GetAnswerByID(uint(answerID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Answer not found")
	}

	return ctx.JSON(http.StatusOK, answer)
}

func (c *AnswerController) UpdateAnswer(ctx echo.Context) error {
	id := ctx.Param("id")

	answerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid answer ID")
	}

	answer := new(models.Answer)
	if err := ctx.Bind(answer); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	answer.ID = uint(answerID)

	err = c.useCase.UpdateAnswer(answer)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update answer")
	}

	return ctx.JSON(http.StatusOK, answer)
}

func (c *AnswerController) DeleteAnswer(ctx echo.Context) error {
	id := ctx.Param("id")

	answerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid answer ID")
	}

	answer := new(models.Answer)
	answer.ID = uint(answerID)

	err = c.useCase.DeleteAnswer(answer)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete answer")
	}

	return ctx.NoContent(http.StatusNoContent)
}
