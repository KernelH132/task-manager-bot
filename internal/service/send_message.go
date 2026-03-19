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
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", token)

	// 1. Prepare the "Buffer" (the body of the request)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 2. Add the text fields
	writer.WriteField("chat_id", strconv.FormatInt(chatID, 10))
	writer.WriteField("caption", caption)
	writer.WriteField("parse_mode", "Markdown")

	// 3. Add the file
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("could not open image: %w", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("photo", imagePath)
	if err != nil {
		return err
	}
	io.Copy(part, file)

	// 4. IMPORTANT: Close the writer BEFORE creating the request
	writer.Close()

	// 5. Create and send the request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	// Set the specific Multipart content type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram error: %s", resp.Status)
	}

	return nil
}
