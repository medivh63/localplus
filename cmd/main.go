package main

import (
	"log"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2/v4"
	"localplus/internal/handlers"
	"localplus/internal/repository"
	"localplus/internal/service"
)

func main() {
	r := gin.Default()

	// 设置模板引擎
	r.HTMLRender = newPongoRenderer("./templates")

	// 设置静态文件路径
	r.Static("/static", "./static")

	// 初始化数据库
	db, err := sql.Open("sqlite3", "quiz.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	questionRepo := sqlite.NewSQLiteQuestionRepository(db)
	quizRepo := sqlite.NewSQLiteQuizRepository(db)

	// 初始化服务和处理器
	quizService := service.NewQuizService(questionRepo, quizRepo)
	quizHandler := handlers.NewQuizHandler(quizService)

	// 设置路由
	r.GET("/", quizHandler.ShowQuiz)
	r.POST("/submit", quizHandler.SubmitQuiz)

	// 启动服务器
	r.Run(":8080")
}

func newPongoRenderer(templatesDir string) gin.HTMLRenderer {
	// 实现pongo2渲染器
	// ...
}
