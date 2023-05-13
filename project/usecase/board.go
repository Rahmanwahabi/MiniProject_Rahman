package usecase

import (
	"MiniProject_Rahman/project/lib/database"
	"MiniProject_Rahman/project/models"
	"errors"
	"fmt"
)

type BoardUsecase interface {
	CreateBoard(board *models.Board) error
	GetBoard(id uint) (board models.Board, err error)
	GetListBoards() (boards []models.Board, err error)
	UpdateBoard(board *models.Board) (err error)
	DeleteBoard(id uint) (err error)
}

type boardUsecase struct {
	boardRepository database.BoardRepository
	taskRepository  database.TaskRepository
}

func NewBoardUsecase(boardRepo database.BoardRepository) *boardUsecase {
	return &boardUsecase{boardRepository: boardRepo}
}

func (b *boardUsecase) CreateBoard(board *models.Board) error {

	// check name cannot be empty
	if board.BoardName == "" {
		return errors.New("board title cannot be empty")
	}

	// check email
	if board.BoardDescription == "" {
		return errors.New("board decription cannot be empty")
	}

	err := b.boardRepository.CreateBoard(board)
	if err != nil {
		return err
	}

	return nil
}

func (b *boardUsecase) GetBoard(id uint) (board models.Board, err error) {
	board, err = b.boardRepository.GetBoard(id)
	if err != nil {
		fmt.Println("GetBoard: Error getting board from database")
		return
	}
	return
}

func (b *boardUsecase) GetListBoards() (boards []models.Board, err error) {
	boards, err = b.boardRepository.GetBoards()
	if err != nil {
		fmt.Println("GetListBoards: Error getting boards from database")
		return
	}
	return
}

func (b *boardUsecase) UpdateBoard(board *models.Board) (err error) {
	err = b.boardRepository.UpdateBoard(board)
	if err != nil {
		fmt.Println("UpdateBoard : Error updating board, err: ", err)
		return
	}

	return
}

func (b *boardUsecase) DeleteBoard(id uint) (err error) {
	board := models.Board{}
	board.ID = id
	err = b.boardRepository.DeleteBoard(&board)
	if err != nil {
		fmt.Println("DeleteBoard : error deleting board, err: ", err)
		return
	}

	return
}
