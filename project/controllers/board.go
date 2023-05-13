package controllers

import (
	"MiniProject_Rahman/project/lib/database"
	"MiniProject_Rahman/project/models"
	"MiniProject_Rahman/project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BoardController interface {
	GetBoardsController(c echo.Context) error
	GetBoardController(c echo.Context) error
	CreateBoardController(c echo.Context) error
	DeleteBoardController(c echo.Context) error
	UpdateBoardController(c echo.Context) error
}

type boardController struct {
	boardUsecase    usecase.BoardUsecase
	boardRepository database.BoardRepository
}

func NewBoardController(
	boardUsecase usecase.BoardUsecase,
	boardRepository database.BoardRepository,
) *boardController {
	return &boardController{
		boardUsecase,
		boardRepository,
	}
}

func (b *boardController) GetBoardsController(c echo.Context) error {
	boards, e := b.boardRepository.GetBoards()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"boards": boards,
	})
}

func (b *boardController) GetBoardController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	board, err := b.boardRepository.GetBoardWithTask(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get board",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"boards": board,
	})
}

// create new board
func (b *boardController) CreateBoardController(c echo.Context) error {
	board := models.Board{}
	c.Bind(&board)

	if err := b.boardUsecase.CreateBoard(&board); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create board",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new board",
		"board":   board,
	})
}

// delete board by id
func (b *boardController) DeleteBoardController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := b.boardUsecase.DeleteBoard(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete board",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf board tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete board",
	})
}

// update board by id
func (b *boardController) UpdateBoardController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	board := models.Board{}
	c.Bind(&board)
	board.ID = uint(id)
	if err := b.boardUsecase.UpdateBoard(&board); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update board",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf board tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update board",
	})
}
