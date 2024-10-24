package service

import (
	"localplus/internal/models"
	"localplus/internal/repository"
	"slices"

	"golang.org/x/exp/rand"
)

type QuizService struct {
	repo *repository.SQLiteRepository
}

func NewQuizService(repo *repository.SQLiteRepository) *QuizService {
	return &QuizService{repo: repo}
}

func (s *QuizService) GetRandomQuestion(quizID string) (*models.Question, error) {
	// 获取已回答的题目id
	answeredIDs, err := s.repo.GetAnsweredQuestionIDs(quizID)
	if err != nil {
		return nil, err
	}
	// 获取所有题目id
	allQuestionIDs := models.GetAllQuestionIDs()
	// 取两个集合的差集
	diffIDs := make([]string, 0)
	for _, id := range allQuestionIDs {
		if !slices.Contains(answeredIDs, id) {
			diffIDs = append(diffIDs, id)
		}
	}
	if len(diffIDs) > 0 {
		// 从差集中随机获取一个id
		randomID := diffIDs[rand.Intn(len(diffIDs))]
		// 根据id查询数据库获取题目
		question, err := s.repo.GetQuestionByID(randomID)
		if err != nil {
			return nil, err
		}
		return question, nil
	} else {
		// 处理 diffIDs 为空的情况
		// 例如，返回一个错误或使用默认值
		return nil, models.ErrQuestionNotFound
	}
}

func (s *QuizService) SaveQuizAnswer(quiz *models.Quiz) error {
	return s.repo.SaveQuizAnswer(quiz)
}

func (s *QuizService) GetQuizResult(quizID string) ([]models.Quiz, error) {
	return s.repo.GetQuizResult(quizID)
}
