package service

import (
	"database/sql"
	"fmt"
)

// SaveTaskToDB ensures the user exists and then saves their task
func SaveUsernameToDB(db *sql.DB, chatID int64, username string) error {
	// 1. First, we "Upsert" the user to get their internal BIGSERIAL ID.
	// This ensures that even if they haven't "registered," they exist now.
	var internalUserID int64
	userQuery := `
        INSERT INTO users (chat_id) 
        VALUES ($1) 
        ON CONFLICT (chat_id) DO UPDATE SET updated_at = NOW()
        RETURNING id;
    `
	err := db.QueryRow(userQuery, chatID).Scan(&internalUserID)
	if err != nil {
		return fmt.Errorf("could not sync user: %v", err)
	}

	// 2. Now we insert the task using that internal ID.
	taskQuery := `
        INSERT INTO tasks (user_id, username) 
        VALUES ($1, $2);
    `
	_, err = db.Exec(taskQuery, internalUserID, username)
	if err != nil {
		return fmt.Errorf("could not save task: %v", err)
	}

	return nil
}
