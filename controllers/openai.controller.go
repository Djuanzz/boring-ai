package controllers

import (
	"net/http"

	"github.com/Djuanzz/boring-ai/services"

	"github.com/gin-gonic/gin"
)

type OpenAIController struct {
	openAIService services.OpenAIService
}

// Constructor
func NewOpenAIController(openAIService services.OpenAIService) *OpenAIController {
	return &OpenAIController{openAIService: openAIService}
}

// Handler untuk menerima prompt dari user dan mengirimkannya ke OpenRouter
func (o *OpenAIController) GenerateResponse(c *gin.Context) {
	var request struct {
		Prompt string `json:"prompt" binding:"required"`
	}

	// Parse request body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prompt is required"})
		return
	}

	// Panggil service untuk mendapatkan response AI
	response, err := o.openAIService.GenerateResponse(request.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kirim response ke user
	c.JSON(http.StatusOK, gin.H{"response": response})
}
