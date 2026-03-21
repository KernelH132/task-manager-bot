// Package service contains the business logic for the bot, including functions to send messages, photos, and chat actions to users via the Telegram Bot API. It also includes functions to manage user states and profiles, as well as handling main menu interactions.
package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/KernelH132/ryuk-bot/internal/llm"
	"github.com/KernelH132/ryuk-bot/internal/models"
)

// SendMessage sends a text message to a specified chat ID
func SendMessage(ctx context.Context, chatID int64, message string) error {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return errors.New("telegram bot token not set in environment")
	}

	reqBody := &models.SendMessageReqBody{
		ChatID:    chatID,
		Text:      message,
		ParseMode: "Markdown",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", res.Status)
	}

	return nil
}

// SendPhotoWithCaption sends a photo with an optional caption to a specified chat ID
func SendPhotoWithCaption(ctx context.Context, chatID int64, photoURL string, caption string) error {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return errors.New("telegram bot token not set in environment")
	}

	reqBody := &models.SendPhotoReqBody{
		ChatID:  chatID,
		Photo:   photoURL,
		Caption: caption,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", token)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", res.Status)
	}

	return nil
}

// SendChatAction sends a chat action (like "typing") to a specified chat ID
func SendChatAction(ctx context.Context, chatID int64, action string) {
	token := os.Getenv("BOT_TOKEN")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendChatAction?chat_id=%d&action=%s", token, chatID, action)
	_, _ = http.Get(url)
}

// sendrandomquote sends a random motivational quote to the user
func SendRandomQuote(ctx context.Context, chatID int64) {
	quotes := []string{
		"Keep going, you're closer than you think.",
		"Small steps still move you forward.",
		"Progress beats perfection.",
		"You've handled worse, you'll handle this.",
		"Start now, figure it out later.",
		"Discipline creates freedom.",
		"Your future is built today.",
		"Stay consistent, results will come.",
		"Hard things build strong people.",
		"Do it tired, do it unmotivated, just do it.",
		"Every day is a new chance to improve.",
		"Focus on what you can control.",
		"You don’t need permission to grow.",
		"Action removes doubt.",
		"Keep showing up, that’s the secret.",
		"Greatness starts with small habits.",
		"Your effort is never wasted.",
		"Make today count.",
		"Comfort zones don’t build success.",
		"Be better than yesterday.",
		"You are capable of more.",
		"Stay patient, stay working.",
		"Momentum comes from movement.",
		"Nothing changes if nothing changes.",
		"Push through, you're almost there.",
		"Believe it, then build it.",
		"Growth feels uncomfortable for a reason.",
		"Win the day, one task at a time.",
		"Your mindset shapes your reality.",
		"Done is better than perfect.",
		"Chase progress, not approval.",
		"Turn effort into results.",
		"Don’t quit before it works.",
		"Make yourself proud.",
		"Stay focused, ignore distractions.",
		"You’re building something bigger.",
		"Energy flows where focus goes.",
		"Level up quietly.",
		"Consistency beats talent.",
		"Work now, shine later.",
		"Keep it simple, keep it moving.",
		"Results come to those who persist.",
		"Trust the process.",
		"Pressure creates diamonds.",
		"Do the work, reap the reward.",
		"Your limits are not fixed.",
		"Stay hungry, stay driven.",
		"Every effort counts.",
		"Build, learn, repeat.",
		"Success starts with trying.",
	}

	randomIndex := rand.Intn(len(quotes))
	randomQuote := quotes[randomIndex]

	SendMessage(ctx, chatID, randomQuote)
}

func HandleAIRequest(ctx context.Context, chatID int64, input string, llmSvc *llm.LLMService) {
	SendChatAction(ctx, chatID, "typing")
	SendChatAction(ctx, chatID, "typing")
	response, err := llmSvc.Generate(input)
	if err != nil {
		fmt.Println("LLM Error:", err)
		SendMessage(ctx, chatID, "Sorry, I'm having trouble thinking right now. 😵‍💫")
		return
	}

	SendMessage(ctx, chatID, response)
}
