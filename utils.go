package tegola

import (
	"bytes"
	"net/http"
	"time"
)

var TelegramBotApiUrl string = "https://api.telegram.org/bot"

// UnixToHumanDate returns the unix timestamp date in a readable format
func (m *Message) UnixToHumanDate(timezone string) (string, error) {
	var (
		tz       *time.Location
		datetime string
		err      error
	)
	tz, err = time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	datetime = time.Unix(m.Date, 0).In(tz).String()
	return datetime, err
}

func (b *Bot) Get(method string) (*http.Response, error) {
	return b.Client.Get(TelegramBotApiUrl + b.BotToken + method)
}

func (b *Bot) Post(method string, payload []byte) (*http.Response, error) {
	return b.Client.Post(TelegramBotApiUrl+b.BotToken+method, "application/json", bytes.NewBuffer(payload))
}

// IsAllowedChat checks if the incoming chat is allowed by the user rules
func (b *Bot) IsAllowedChat(chatId int64) bool {
	for _, id := range b.AllowedChats {
		if chatId == id {
			return true
		}
	}
	return false
}

// IsAllowedUser checks if the incoming chat user is allowed by the user rules
func (b *Bot) IsAllowedUser(from User) bool {
	for _, user := range b.AllowedUsers {
		if from.Id == user.Id && from.Username == user.Username {
			return true
		}
	}
	return false
}

// AllowedMessage checks if the incoming chat message is allowed by the user rules
func (b *Bot) AllowedMessage(m Message) bool {
	allowed := b.IsAllowedChat(m.Chat.Id) && b.IsAllowedUser(m.From)
	return allowed
}

// Debug writes debug messages to the terminal
func (b *Bot) Debug() {

}

// Echo repeats last user message back to the chat
func (b *Bot) Echo() {

}

// EchoDebug debugs the bot back to the chat
func (b *Bot) EchoDebug() {

}

// GhostDebug debugs the bot to the admin chat
func (b *Bot) GhostDebug() {

}

// ParentalControl asks admin for permission to talk to strangers
func (b *Bot) ParentalControl() {

}
