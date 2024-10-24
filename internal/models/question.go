package models

import "time"

type Question struct {
	QuestionID string `json:"question_id" gorm:"primaryKey"`
	Content    string `json:"content"`
	Options    string `json:"options"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

// TableName 指定表名
func (Question) TableName() string {
	return "questions"
}