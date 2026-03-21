package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type LLMService struct {
	APIKey string
}

func New() *LLMService {
	return &LLMService{
		APIKey: os.Getenv("OPENAI_API_KEY"),
	}
}

// OpenAI style message structure
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type requestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func (s *LLMService) Generate(prompt string) (string, error) {
	// request body for OpenAI
	body := requestBody{
		Model: "gpt-4o-mini", // Correct model name
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// 2. OpenAI Endpoint
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+s.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from model")
}
