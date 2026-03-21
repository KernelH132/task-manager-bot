package models

type SendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type SendPhotoReqBody struct {
	ChatID    int64  `json:"chat_id"`
	Photo     string `json:"photo"`
	Caption   string `json:"caption,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
}

type chat struct {
	ID int64 `json:"id"`
}

type message struct {
	Text string `json:"text"`
	Chat chat   `json:"chat"`
}

type WebhookReqBody struct {
	Message message `json:"message"`
}
