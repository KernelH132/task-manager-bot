package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/KernelH132/pingme/internal/repository"
)

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
