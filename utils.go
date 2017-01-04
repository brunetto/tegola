package tegola

import (
	"bytes"
	"net/http"
	"time"
	"github.com/brunetto/goutils/debug"
	"errors"
	"strconv"
	"io/ioutil"
	"log"
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

func CheckHttpErrors (resp *http.Response, url string) (*http.Response, error) {
	var (
		respBytes []byte
		respString string
		err error
	)
	if resp.StatusCode != 200 {
		respBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("CheckHttpErrors: can't read http response")
		}
		respString = string(respBytes)
		return resp, errors.New("HTTP error : " +
					strconv.Itoa(resp.StatusCode) +
					"\n" + respString +
					"\n URL: " + url)
	}
	return resp, nil
}

func (b *Bot) Get(method string) (*http.Response, error) {
	var (
		url string
		resp *http.Response
		err error
	)
	url = TelegramBotApiUrl + b.BotToken + "/" + method
	debug.LogDebug(b.Debug, "Bot get url: ", url)
	resp, err = b.Client.Get(url)
	if err != nil {
		return resp, err
	}
	resp, err = CheckHttpErrors(resp, url)
	return resp, err
}

func (b *Bot) Post(method string, payload []byte) (*http.Response, error) {
	var (
		url string
		resp *http.Response
		err error
	)
	url = TelegramBotApiUrl + b.BotToken + "/" + method
	debug.LogDebug(b.Debug, "Bot post url: ", url)
	resp, err = b.Client.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return resp, err
	}
	resp, err = CheckHttpErrors(resp, url)
	return resp, err
}

// IsAllowedChat checks if the incoming chat is allowed by the user rules
func (b *Bot) IsAllowedChat(chatId int64) bool {
	for _, id := range b.AllowedChats {
		if chatId == id {
			debug.LogDebug(b.Debug, "Allowed chat ", chatId)
			log.Println(b.Debug)
			return true
		}
	}
	debug.LogDebug(b.Debug, "Forbidden chat ", chatId)
	return false
}

// IsAllowedUser checks if the incoming chat user is allowed by the user rules
func (b *Bot) IsAllowedUser(from User) bool {
	for _, user := range b.AllowedUsers {
		if from.Id == user.Id && from.Username == user.Username {
			debug.LogDebug(b.Debug, "Allowed user ", from)
			return true
		}
	}
	debug.LogDebug(b.Debug, "Forbidden user ", from)
	return false
}

// AllowedMessage checks if the incoming chat message is allowed by the user rules
func (b *Bot) AllowedMessage(m Message) bool {
	allowed := b.IsAllowedChat(m.Chat.Id) && b.IsAllowedUser(m.From)
	return allowed
}

// Debug writes debug messages to the terminal
//func (b *Bot) Debug() {
//
//}

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
