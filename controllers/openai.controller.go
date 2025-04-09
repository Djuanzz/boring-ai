package controllers

import (
	"net/http"

	"github.com/Djuanzz/boring-ai/dto"
	"github.com/Djuanzz/boring-ai/services"

	"github.com/gin-gonic/gin"
)

type OpenAIController interface {
	GenerateResponse(c *gin.Context)
}

type openAIController struct {
	openAIService services.OpenAIService
}

// Constructor
func NewOpenAIController(openAIService services.OpenAIService) OpenAIController {
	return &openAIController{openAIService: openAIService}
}

// Handler untuk menerima prompt dari user dan mengirimkannya ke OpenRouter
func (o *openAIController) GenerateResponse(c *gin.Context) {

	var req dto.OpenAiRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prompt is required"})
		return
	}

	// Panggil service untuk mendapatkan response AI
	response, err := o.openAIService.GenerateResponse(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kirim response ke user
	c.JSON(http.StatusOK, gin.H{"response": response})
}
