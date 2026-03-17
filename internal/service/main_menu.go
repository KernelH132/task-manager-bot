package service

import (
	"strings"

	"github.com/KernelH132/pingme/internal/repository"
)

func HandleMainMenu(chatID int64, input string) {
	if strings.ToLower(input) == "/register" {
		// Step 1: Tell the DB we are waiting for a task
		SetUserState(repository.DB, chatID, "awaiting_username")

		// Step 2: Ask the user for the info
		SendMessage(chatID, "Please enter a username😊.")
	}
}

func HandleUsernameCreation(chatID int64, username string) {
	// Step 1: Save the task
	err := SaveUsernameToDB(repository.DB, chatID, username)
	if err != nil {
		SendMessage(chatID, "System error. Try again later.")
		return
	}

	// Step 2: RESET to idle so they can use other commands again
	SetUserState(repository.DB, chatID, "idle")

	// Step 3: Success message
	SendMessage(chatID, "✅ Task saved!")

}
