package service

import (
	"your-username/alberta-class7-quiz/internal/models"
	"your-username/alberta-class7-quiz/internal/repository"
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
