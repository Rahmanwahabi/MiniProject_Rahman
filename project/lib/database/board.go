package database

import (
	"MiniProject_Rahman/project/config"
	"MiniProject_Rahman/project/models"

	"gorm.io/gorm"
)

type BoardRepository interface {
	CreateBoard(board *models.Board) error
	GetBoards() (boards []models.Board, err error)
	GetBoard(id uint) (board models.Board, err error)
	GetBoardsByUserId(userID uint) (boards []models.Board, err error)
	GetBoardWithTask(id uint) (board models.Board, err error)
	UpdateBoard(board *models.Board) error
	DeleteBoard(board *models.Board) error
}

type boardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) *boardRepository {
	return &boardRepository{db}
}

func (b *boardRepository) CreateBoard(board *models.Board) error {

	if err := config.DB.Create(board).Error; err != nil {
		return err
	}
	return nil

}

func (b *boardRepository) GetBoards() (boards []models.Board, err error) {
	if err = config.DB.Model(&models.Board{}).Preload("Tasks").Find(&boards).Error; err != nil {
		return
	}
	return
}

func (b *boardRepository) GetBoard(id uint) (board models.Board, err error) {
	board.ID = id
	if err = config.DB.First(&board).Error; err != nil {
		return
	}
	return
}

func (b *boardRepository) GetBoardWithTask(id uint) (board models.Board, err error) {
	board.ID = id
	if err = config.DB.Model(&models.Board{}).Preload("Tasks").First(&board).Error; err != nil {
		return
	}
	return
}

func (b *boardRepository) GetBoardsByUserId(userID uint) (boards []models.Board, err error) {
	if err = config.DB.Where("user_id = ?", userID).Find(&boards).Error; err != nil {
		return
	}
	return
}

func (b *boardRepository) UpdateBoard(board *models.Board) error {
	if err := config.DB.Updates(board).Error; err != nil {
		return err
	}
	return nil
}

func (b *boardRepository) DeleteBoard(board *models.Board) error {
	if err := config.DB.Delete(board).Error; err != nil {
		return err
	}
	return nil
}
