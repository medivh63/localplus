package service

import (
	"localplus/internal/models"
	"localplus/internal/repository"
)

type QuizService struct {
	repo *repository.SQLiteRepository
}

func NewQuizService(repo *repository.SQLiteRepository) *QuizService {
	return &QuizService{repo: repo}
}

func (s *QuizService) GetRandomQuestion() (*models.Question, error) {
	return s.repo.GetRandomQuestion()
}

func (s *QuizService) SaveQuizAnswer(quiz *models.Quiz) error {
	return s.repo.SaveQuizAnswer(quiz)
}

func (s *QuizService) GetQuizResult(quizID string) ([]models.Quiz, error) {
	return s.repo.GetQuizResult(quizID)
}
