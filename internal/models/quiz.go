package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	QuizID     string    `json:"quiz_id"`
	QuestionID string    `json:"question_id"`
	IsCorrect  bool      `json:"is_correct"`
}

// TableName 指定表名
func (Quiz) TableName() string {
	return "quizzes"
}
