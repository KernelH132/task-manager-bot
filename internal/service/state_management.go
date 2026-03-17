package service

import (
	"database/sql"
	"fmt"

	"github.com/KernelH132/pingme/internal/repository"
)

// GetUserState retrieves the current string state for a chat_id
func GetUserState(db *sql.DB, chatID int64) string {
	var state string
	query := `SELECT current_state FROM users WHERE chat_id = $1`

	err := repository.DB.QueryRow(query, chatID).Scan(&state)
	if err != nil {
		if err == sql.ErrNoRows {
			return "idle" // Default for new users
		}
		fmt.Println("Error fetching state:", err)
		return "idle"
	}
	return state
}

// SetUserState updates the user's state in the database
func SetUserState(db *sql.DB, chatID int64, newState string) error {
	query := `
        INSERT INTO users (chat_id, current_state)
        VALUES ($1, $2)
        ON CONFLICT (chat_id) 
        DO UPDATE SET current_state = EXCLUDED.current_state, updated_at = NOW();
    `
	_, err := repository.DB.Exec(query, chatID, newState)
	return err
}
