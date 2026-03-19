package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/KernelH132/pingme/internal/models"
)

func SendMessage(chatID int64, message string) error {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return errors.New("telegram bot token not set in environment")
	}

	reqBody := &models.SendMessageReqBody{
		ChatID: chatID,
		Text:   message,
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
func SendPhotoWithCaption(chatID int64, imagePath string, caption string) error {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return errors.New("telegram bot token not set in environment")
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", token)

	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add form fields
	writer.WriteField("chat_id", strconv.FormatInt(chatID, 10))
	writer.WriteField("caption", caption)
	writer.WriteField("parse_mode", "Markdown")

	// Create the file part
	part, err := writer.CreateFormFile("photo", imagePath)
	if err != nil {
		return err
	}

	if _, err = io.Copy(part, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram api error: %s", res.Status)
	}

	return nil
}
