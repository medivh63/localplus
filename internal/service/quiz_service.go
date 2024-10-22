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

func (s *QuizService) SubmitQuiz(quiz *models.Quiz) error {
	return s.repo.SaveQuizResult(quiz)
}
