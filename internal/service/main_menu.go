package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/KernelH132/ryuk-bot/internal/messages"
	"github.com/KernelH132/ryuk-bot/internal/repository"
)

func HandleMainMenu(ctx context.Context, chatID int64, input string) {
	switch {

	case input == "/start":

		SendChatAction(ctx, chatID, "typing")

		err := SendPhotoWithCaption(ctx, chatID, "https://pin.it/6dAfyFQff", messages.WelcomeMessage)
		if err != nil {
			fmt.Println("Error sending welcome message:", err)
			SendMessage(ctx, chatID, "Welcome to Ryuk Bot! Use /help to see available commands.")
		}

	case strings.ToLower(input) == "/register":

		err := SetUserState(ctx, repository.DB, chatID, "awaiting_username")
		if err != nil {
			SendMessage(ctx, chatID, "Connection error. Please try /register again.")
			return
		}

		SendMessage(ctx, chatID, "Please enter a username 😊.")

	case strings.ToLower(input) == "/help":
		SendMessage(ctx, chatID, messages.HelpMessage)

	case input == "/ping":
		SendMessage(ctx, chatID, messages.Ping)

	case input == "/profile":

		username, err := GetProfile(ctx, chatID)

		if err != nil {
			SendMessage(ctx, chatID, "Error fetching profile. Please register first using /register.")
			return
		}
		SendMessage(ctx, chatID, fmt.Sprintf("Your username is: %s", username))

	case input == "/quote":
		SendRandomQuote(ctx, chatID)

	default:
		SendMessage(ctx, chatID, "Sorry, I didn't understand that command. Use /help to see available commands.")
	}

}
