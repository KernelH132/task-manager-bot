package service

import (
	"context"
	"database/sql"
	"fmt"
)

// SaveTaskToDB ensures the user exists and then saves their task
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
