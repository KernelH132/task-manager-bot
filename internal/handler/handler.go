package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/KernelH132/pingme/internal/models"
	"github.com/KernelH132/pingme/internal/repository"
	"github.com/KernelH132/pingme/internal/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	body := &models.WebhookReqBody{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		fmt.Println("Could not decode request body", err)
		return
	}

	userInput := strings.ToLower(strings.TrimSpace(body.Message.Text))
	chatID := body.Message.Chat.ID

	current_state := service.GetUserState(repository.DB, chatID)

	switch current_state {
	case "awaiting_username":
		service.HandleUsernameCreation(chatID, userInput)
	case "idle":
		service.HandleMainMenu(chatID, userInput)
	default:
		// Fallback to idle if something goes wrong
		service.SetUserState(repository.DB, chatID, "idle")
		service.HandleMainMenu(chatID, userInput)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Response processed successfully")
}
