package tegola

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/brunetto/goutils/debug"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
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

func CheckHttpErrors(resp *http.Response, url string) (*http.Response, error) {
	var (
		respBytes  []byte
		respString string
		err        error
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
		url  string
		resp *http.Response
		err  error
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
		url  string
		resp *http.Response
		err  error
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
	// If no allowed chats are listed, every chat is allowed
	if len(b.AllowedChats) == 0 {
		return true
	}

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
	// If no allowed users are listed, every user is allowed
	if len(b.AllowedUsers) == 0 {
		return true
	}

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

func (b *Bot) Start(updatesHandler func(*Bot, *sync.WaitGroup)) error {
	var (
		wg  sync.WaitGroup
		err error
	)
	debug.LogDebug(b.Debug, "Setting up handler")

	http.HandleFunc(b.ListenRoute, b.WebhooksUpdatesHandler)

	debug.LogDebug(b.Debug, "Listening for updates from webhook")

	wg.Add(1)
	go updatesHandler(b, &wg)

	err = http.ListenAndServe("127.0.0.1:8443", nil)
	close(b.UpdatesChan)
	wg.Wait()
	if err != nil {
		return err
	}
	return nil
}

// Echo repeats last user message back to the chat
func Echo(b *Bot, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		reply Message
	)

	for u := range b.UpdatesChan {
		messageText := u.Message.Text
		chatId := strconv.Itoa(int(u.Message.Chat.Id))
		sender := u.Message.From
		date, err := u.Message.UnixToHumanDate(b.TimeZone)

		if err != nil {
			log.Println(err)
		}

		// Echo to admin
		// @TODO: format message
		replyText := "Echo" + "\n====" +
			"\nSender: " + sender.Username + " - id: " + strconv.Itoa(int(sender.Id)) +
			"\nChat: " + chatId +
			"\nTimestamp " + date +
			"\nUpdate n.: " + strconv.Itoa(int(u.UpdateId)) +
			"\nMessage n.: " + strconv.Itoa(int(u.Message.MessageId)) +
			"\nMessage:\n " + messageText + "\n"

		sp := SendMessagePayload{
			ChatId: u.Message.Chat.Id, /*b.AdminChats[0]*/
			Text:   replyText,
		}

		reply, err = b.SendMessage(sp)
		if err != nil {
			log.Println(err)
		} else {

			fmt.Println("Echoed message sent back  to chat " + strconv.Itoa(int(u.Message.Chat.Id)) + " is: ")
			fmt.Println(reply.Text)
		}
	}
}

// Echo repeats last user message back to the chat
func EchoHandler(b *Bot, c CmdData, u Update) error {
	var (
		reply Message
	)
	messageText := u.Message.Text
	chatId := strconv.Itoa(int(u.Message.Chat.Id))
	sender := u.Message.From
	date, err := u.Message.UnixToHumanDate(b.TimeZone)

	if err != nil {
		log.Println(err)
	}

	// Echo to admin
	// @TODO: format message
	replyText := "Echo" + "\n====" +
		"\nSender: " + sender.Username + " - id: " + strconv.Itoa(int(sender.Id)) +
		"\nChat: " + chatId +
		"\nTimestamp " + date +
		"\nUpdate n.: " + strconv.Itoa(int(u.UpdateId)) +
		"\nMessage n.: " + strconv.Itoa(int(u.Message.MessageId)) +
		"\nMessage:\n " + messageText + "\n"

	if c.Cmd != "" {
		replyText += "Found command " + c.Cmd + "\n"
	} else {
		replyText += "No command found.\n"
	}

	sp := SendMessagePayload{
		ChatId: u.Message.Chat.Id, /*b.AdminChats[0]*/
		Text:   replyText,
	}

	reply, err = b.SendMessage(sp)
	if err != nil {
		log.Println(err)
	} else {

		fmt.Println("Echoed message sent back  to chat " + strconv.Itoa(int(u.Message.Chat.Id)) + " is: ")
		fmt.Println(reply.Text)
	}
	return nil
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

//func (b *Bot) Shutdown() {
//	if b.MongoSession {
//		b.MongoSession.Close()
//	}
//}