package models

import "time"

type Quiz struct {
	ID         string    `json:"id"`
	QuestionID string    `json:"question_id"`
	IsCorrect  bool      `json:"is_correct"`
	CreatedAt  time.Time `json:"created_at"`
}
