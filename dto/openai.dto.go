package dto

type OpenAiRequest struct {
	Prompt string `json:"prompt" form:"prompt" binding:"required"`
}
