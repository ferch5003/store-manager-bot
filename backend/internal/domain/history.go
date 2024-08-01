package domain

type History struct {
	ID          int    `json:"id" db:"id" `
	UserMessage string `json:"user_message" db:"user_message"`
	BotResponse string `json:"bot_response" db:"bot_response"`
	Feedback    bool   `json:"feedback" db:"feedback"`
}
