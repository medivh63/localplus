package handlers

import (
	"net/http"
	"time"

	"localplus/internal/models"
	"localplus/internal/service"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QuizHandler struct {
	quizService *service.QuizService
}

func NewQuizHandler(quizService *service.QuizService) *QuizHandler {
	return &QuizHandler{quizService: quizService}
}

func (h *QuizHandler) Index(c *gin.Context) {
	quizID := uuid.New().String()
	// 如果cookie存在，并且已过期就设置新的cookie
	if cookie, err := c.Request.Cookie("quizID"); err == nil {
		if time.Now().After(cookie.Expires) {
			c.SetCookie("quizID", quizID, -1, "/", "localplus.com", false, true)
		} else {
			quizID = cookie.Value
		}
	}
	data := pongo2.Context{"quizID": quizID}
	RenderTemplate(c, "index.html", data)
}

func (h *QuizHandler) StartQuiz(c *gin.Context) {
	quizID := c.Param("quizID")

	question, err := h.quizService.GetRandomQuestion(quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quiz_id":   quizID,
		"question":  question,
	})
}

func (h *QuizHandler) AnswerQuestion(c *gin.Context) {
	var answer struct {
		QuestionID string `json:"question_id"`
		Answer     string `json:"answer"`
	}

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	quizID := c.Param("quizID")

	// 这里应该检查答案是否正确,但为了简化,我们假设所有答案都是正确的
	quiz := &models.Quiz{
		QuizID:     quizID,
		QuestionID: answer.QuestionID,
		IsCorrect:  true,
	}

	if err := h.quizService.SaveQuizAnswer(quiz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save answer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer saved successfully"})
}

func (h *QuizHandler) GetQuizResult(c *gin.Context) {
	quizID := c.Param("quizID")

	results, err := h.quizService.GetQuizResult(quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get quiz results"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
