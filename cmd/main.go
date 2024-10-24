package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"localplus/internal/handlers"
	"localplus/internal/repository"
	"localplus/internal/service"
	"localplus/internal/models"
)

func main() {
	db, err := repository.NewSQLiteDB("/Users/medivh/local.db")
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&models.Question{}, &models.Quiz{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repo, err := repository.NewSQLiteRepository("/Users/medivh/local.db")
	if err != nil {
		log.Fatalf("无法创建仓库: %v", err)
	}

	quizService := service.NewQuizService(repo)
	quizHandler := handlers.NewQuizHandler(quizService)

	r := gin.Default()

	r.POST("/quiz", quizHandler.StartQuiz)
	r.POST("/quiz/:quizID/answer", quizHandler.AnswerQuestion)
	r.GET("/quiz/:quizID/result", quizHandler.GetQuizResult)

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
