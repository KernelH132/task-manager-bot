package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/KernelH132/ryuk-bot/internal/models"
	"github.com/KernelH132/ryuk-bot/internal/repository"
	"github.com/KernelH132/ryuk-bot/internal/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	body := &models.WebhookReqBody{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		fmt.Println("Could not decode request body", err)
		return
	}

	userInput := strings.ToLower(strings.TrimSpace(body.Message.Text))
	chatID := body.Message.Chat.ID

	current_state := service.GetUserState(ctx, repository.DB, chatID)

	switch current_state {
	case "awaiting_username":
		service.HandleUsernameCreation(ctx, chatID, userInput)
	case "idle":
		service.HandleMainMenu(ctx, chatID, userInput)
	default:
		// Fallback to idle if something goes wrong
		service.SetUserState(ctx, repository.DB, chatID, "idle")
		service.HandleMainMenu(ctx, chatID, userInput)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Response processed successfully")
}
