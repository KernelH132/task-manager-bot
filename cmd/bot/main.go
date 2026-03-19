package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KernelH132/pingme/internal/handler"
	"github.com/KernelH132/pingme/internal/repository"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

}

func setTelegramWebhook() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("telegram bot token not set in environment")
	}
	publicURL := os.Getenv("RAILWAY_PUBLIC_DOMAIN")
	if publicURL == "" {
		log.Fatal("RAILWAY_PUBLIC_DOMAIN environment variable is not set")
	}
	webhookURL := fmt.Sprintf("https://%s/", publicURL)
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", token, webhookURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Failed to set Telegram webhook: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to set Telegram webhook, status code: %d", resp.StatusCode)
	}

	log.Println("Telegram webhook set successfully")
}

func main() {
	repository.Connect()

	go setTelegramWebhook()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on :" + port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Handler)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
