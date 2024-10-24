package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"localplus/internal/models"
	"localplus/internal/service"
)

type QuizHandler struct {
	quizService *service.QuizService
}

func NewQuizHandler(quizService *service.QuizService) *QuizHandler {
	return &QuizHandler{quizService: quizService}
}

func (h *QuizHandler) StartQuiz(c *gin.Context) {
	question, err := h.quizService.GetRandomQuestion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get question"})
		return
	}

	quizID := uuid.New().String()

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
