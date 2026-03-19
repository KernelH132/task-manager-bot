package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/KernelH132/pingme/internal/repository"
)

func HandleMainMenu(ctx context.Context, chatID int64, input string) {
	switch {
	case input == "/start":
		SendMessage(chatID, `
👋 Welcome to 𝚁𝚢𝚞𝚔 𝙱𝚘𝚝!

Get started by using:
/help — view available commands

View main group to connect with other users, get updates and support:

• Main group: https://t.me/ryuk_bott

🟢 Ready
`)

	case strings.ToLower(input) == "/register":
		err := SetUserState(ctx, repository.DB, chatID, "awaiting_username")
		if err != nil {
			SendMessage(chatID, "Connection error. Please try /register again.")
			return
		}

		SendMessage(chatID, "Please enter a username 😊.")

	case strings.ToLower(input) == "/help":
		SendMessage(chatID, `
╭─── ⚡ RYUK BOT ───╮
│
│  👤 /register  → Create/edit profile
│  ❓ /help      → Show menu
│  👤 /profile   → View profile
│  🏓 /ping      → Check if bot is alive
│  🌐 /group     → Join our main group
╰─────────────────╯

• under construction...
• Main group: https://t.me/ryuk_bott

`)
	case input == "/ping":
		SendMessage(chatID, "pong")

	case input == "/group":
		SendMessage(chatID, "Join our main group to connect with other users, get updates and support:\n\nhttps://t.me/ryuk_bott")

	case input == "/profile":
		username, err := GetProfile(ctx, chatID)
		if err != nil {
			SendMessage(chatID, "Error fetching profile. Please register first using /register.")
			return
		}
		SendMessage(chatID, fmt.Sprintf("Your username is: %s", username))
	}

}

func HandleUsernameCreation(ctx context.Context, chatID int64, username string) {
	// Step 1: Save the username to the database
	err := SaveUsernameToDB(ctx, repository.DB, chatID, username)
	if len(username) >= 32 {
		SendMessage(chatID, "That username is too long! Keep it under 32 characters.")
		return

	}
	if err != nil {
		SendMessage(chatID, "System error. Try again later.")
		return
	}

	// Step 2: RESET to idle so they can use other commands again
	err = SetUserState(ctx, repository.DB, chatID, "idle")
	if err != nil {
		fmt.Printf("Warning: failed to reset user state for %d: %v\n", chatID, err)
		return
	}

	createdMessage := fmt.Sprintf("Hi %s, your username has been created!🚀", username)

	// Step 3: Success message
	SendMessage(chatID, createdMessage)

}
