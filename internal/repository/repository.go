package repository

import "localplus/internal/models"

type QuestionRepository interface {
    GetRandomQuestion() (*models.Question, error)
    GetQuestionByID(id string) (*models.Question, error)
    // 其他问题相关方法...
}

type QuizRepository interface {
    SaveQuizResult(quiz *models.Quiz) error
    GetQuizResults(userID string) ([]*models.Quiz, error)
    // 其他测验相关方法...
}

// 可能还有其他repository接口...
