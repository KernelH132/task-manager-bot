// This file contains all the user-related service functions, such as fetching and saving user profiles, handling username creation flow, and any other user-specific logic.
package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/KernelH132/ryuk-bot/internal/repository"
)

// get user profile information from the database
func GetProfile(ctx context.Context, chatID int64) (string, error) {
	query := `SELECT username FROM users WHERE chat_id = $1`
	var username string
	err := repository.DB.QueryRowContext(ctx, query, chatID).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no profile found for chat_id %d", chatID)
		}
		return "", fmt.Errorf("error fetching profile: %w", err)
	}

	return username, nil
}

// save the user's profile information to the database
func SaveUsernameToDB(ctx context.Context, db *sql.DB, chatID int64, username string) error {
	query := `
        UPDATE users 
        SET username = $1, 
            updated_at = NOW() 
        WHERE chat_id = $2;
    `

	result, err := db.ExecContext(ctx, query, username, chatID)
	if err != nil {
		return fmt.Errorf("could not update username: %w", err)
	}

	// Checking RowsAffected is good practice to ensure the user actually exists
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no user found with chat_id %d", chatID)
	}

	return nil
}

// handle username creation flow, including validation and database saving
func HandleUsernameCreation(ctx context.Context, chatID int64, username string) {

	if len(username) >= 32 {
		SendMessage(ctx, chatID, "That username is too long! Keep it under 32 characters.")
		return

	}
	err := SaveUsernameToDB(ctx, repository.DB, chatID, username)
	if err != nil {
		SendMessage(ctx, chatID, "System error. Try again later.")
		return
	}

	err = SetUserState(ctx, repository.DB, chatID, "idle")
	if err != nil {
		fmt.Printf("Warning: failed to reset user state for %d: %v\n", chatID, err)
		return
	}

	createdMessage := fmt.Sprintf("Hi %s, your username has been created!🚀", username)

	SendMessage(ctx, chatID, createdMessage)

}
