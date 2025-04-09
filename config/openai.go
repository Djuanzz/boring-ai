package config

import "github.com/sashabaranov/go-openai"

func NewOpenAIClient(apiKey string) *openai.Client {
	clientCfg := openai.DefaultConfig(apiKey)

	cli := openai.NewClientWithConfig(clientCfg)

	if cli == nil {
		panic("Error creating OpenAI cli")
	}
	return cli
}
