package models

import (
	"time"
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	QuizID     string    `json:"quiz_id" gorm:"primaryKey"`
	QuestionID string    `json:"question_id"`
	IsCorrect  bool      `json:"is_correct"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

// TableName 指定表名
func (Quiz) TableName() string {
	return "quizzes"
}
