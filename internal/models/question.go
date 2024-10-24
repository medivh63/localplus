package models

import "time"

// 添加全局变量
var allQuestionIDs []string

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

// 添加初始化函数
func InitializeQuestionIDs(ids []string) {
	allQuestionIDs = make([]string, len(ids))
	copy(allQuestionIDs, ids)
}

// 添加获取所有QuestionID的函数
func GetAllQuestionIDs() []string {
	return allQuestionIDs
}
