package usecase

import (
	"MiniProject_Rahman/lib/database"
	"MiniProject_Rahman/models"
)

type AnswerUseCase struct{}

func NewAnswerUseCase() *AnswerUseCase {
	return &AnswerUseCase{}
}

func (uc *AnswerUseCase) CreateAnswer(answer *models.Answer) error {
	err := database.CreateAnswer(answer)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *AnswerUseCase) GetAnswerByID(id uint) (*models.Answer, error) {
	answer, err := database.GetAnswerByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}
	return answer, nil
}

func (uc *AnswerUseCase) UpdateAnswer(answer *models.Answer) error {
	err := database.UpdateAnswer(answer)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *AnswerUseCase) DeleteAnswer(answer *models.Answer) error {
	err := database.DeleteAnswer(answer)
	if err != nil {
		// Handle error
	}
	return nil
}

func (uc *AnswerUseCase) GetAnswers() ([]models.Answer, error) {
	answers, err := database.GetAnswers()
	if err != nil {
		// Handle error (e.g., logging, error wrapping, etc.)
	}
	return answers, err
}
