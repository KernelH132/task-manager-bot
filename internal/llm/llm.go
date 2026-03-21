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

		APIKey: os.Getenv("OPENROUTER_API_KEY"),
	}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type requestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func (s *LLMService) Generate(prompt string) (string, error) {
	body := requestBody{

		Model: "google/gemini-2.0-flash-lite-001",
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// CHANGED: OpenRouter Endpoint
	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+s.APIKey)
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("HTTP-Referer", "https://github.com/KernelH132/ryuk-bot")
	req.Header.Set("X-OpenRouter-Title", "Ryuk Bot")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("OpenRouter error: status %d", resp.StatusCode)
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
