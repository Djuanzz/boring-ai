package services

import (
	"context"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIService interface {
	GenerateResponse(prompt string) (string, error)
}

type openAIService struct {
	client *openai.Client
}

// Constructor untuk OpenAIService dengan OpenRouter
func NewOpenAIService(client *openai.Client) OpenAIService {
	return &openAIService{client: client}
}

// Fungsi untuk memanggil OpenRouter API
func (o *openAIService) GenerateResponse(prompt string) (string, error) {
	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleUser, Content: prompt},
			},
		},
	)

	if err != nil {
		log.Printf("Error calling OpenRouter API: %v", err)
		return "", err
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", nil
}
