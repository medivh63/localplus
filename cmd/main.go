package main

import (
	"log"
	"os"
	"path/filepath"

	"localplus/internal/handlers"
	"localplus/internal/models"
	"localplus/internal/repository"
	"localplus/internal/service"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
)

func loadTemplates() map[string]*pongo2.Template {
	templates := make(map[string]*pongo2.Template)
	templateDir := "templates" // 指定模板目录
	// 遍历模板目录
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			// 加载模板
			tpl, err := pongo2.FromFile(path)
			if err != nil {
				return err
			}
			// 使用相对路径作为键
			relPath, _ := filepath.Rel(templateDir, path)
			templates[relPath] = tpl
		}
		return nil
	})

	log.Printf("load templates success %d", len(templates))
	if err != nil {
		log.Fatalf("加载模板失败: %v", err)
	}

	return templates
}

func main() {
	log.Println("start load templates")
	templates := loadTemplates()
	handlers.InitTemplates(templates)

	log.Println("start connect to database")
	db, err := repository.NewSQLiteDB("/Users/medivh/local.db")
	if err != nil {
		log.Fatalf("can not connect to database: %v", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&models.Question{}, &models.Quiz{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 创建repository
	repo, err := repository.NewSQLiteRepository("/Users/medivh/local.db")
	if err != nil {
		log.Fatalf("can not create repository: %v", err)
	}
	
	// 获取所有question id
	ids, err := repo.GetAllQuestionIDs()
	if err != nil {
		log.Fatalf("can not get all question ids: %v", err)
	}
	models.InitializeQuestionIDs(ids)
	log.Println("query", len(ids), "question ids")

	quizService := service.NewQuizService(repo)
	quizHandler := handlers.NewQuizHandler(quizService)

	r := gin.Default()

	r.GET("/class7", quizHandler.Index)
	r.GET("/class7/quiz/:quizID", quizHandler.StartQuiz)
	r.POST("/class7/quiz/:quizID/answer", quizHandler.AnswerQuestion)
	r.GET("/class7/quiz/:quizID/result", quizHandler.GetQuizResult)

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("server start at port 3000")
}
