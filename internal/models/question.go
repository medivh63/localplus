package models

import (
	"gorm.io/gorm"
)

// 添加全局变量
var allQuestionIDs []string

type Option struct {
	Content   string `json:"content"`
	IsCorrect bool   `json:"is_correct" gorm:"type:numeric"`
}

type Question struct {
	gorm.Model
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	QuestionID string    `json:"question_id"`
	Content    string    `json:"content"`
	Options    []Option  `json:"options" gorm:"type:json"`
	Type       string    `json:"type"`
	Images     string    `json:"image"`
}

// TableName 指定表名
func (Question) TableName() string {
	return "questions"
}

// 添加初始化函数
func InitializeQuestionIDs(ids []string) {
	allQuestionIDs = make([]string, len(ids))
	copy(allQuestionIDs, ids)
}

// 添加获取所有QuestionID的函数
func GetAllQuestionIDs() []string {
	return allQuestionIDs
}
