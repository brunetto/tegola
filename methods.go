package tegola

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func (b *Bot) GetUpdates() ([]Update, []Update, error) {
	var (
		client    = &http.Client{Timeout: time.Second * 10}
		resp      *http.Response
		err       error
		u         Updates
		allowed   = []Update{}
		forbidden = []Update{}
	)
	resp, err = client.Get("https://api.telegram.org/bot" + b.BotToken + "/getUpdates")
	defer resp.Body.Close()
	if err != nil {
		return allowed, forbidden, err
	}

	u = Updates{}
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return allowed, forbidden, err
	}

	for _, Update := range u.Updates {
		if b.AllowedMessage(Update.Message) {
			allowed = append(allowed, Update)
		} else {
			forbidden = append(forbidden, Update)
		}
	}
	return allowed, forbidden, err
}

func (b *Bot) SendSimpleMessage(msgText string) (Message, error) {
	var (
		msgReq = SendMessagePayload{ChatId: b.ChatId, Text: msgText}
	)
	return b.SendMessage(msgReq)
}

func (b *Bot) SendMessage(msgReq SendMessagePayload) (Message, error) {
	var (
		client = &http.Client{Timeout: time.Second * 10}
		resp   *http.Response
		err    error
		m      Message
		msg    []byte
	)

	m = Message{}

	msg, err = json.Marshal(msgReq)
	if err != nil {
		return m, err
	}

	resp, err = client.Post("https://api.telegram.org/bot"+b.BotToken+"/sendMessage", "application/json", bytes.NewBuffer(msg))
	if err != nil {
		return m, err
	}

	err = json.NewDecoder(resp.Body).Decode(&m)
	return m, err
}

func (b *Bot) SetWebhook() {

}

func (b *Bot) GetMe()                  {}
func (b *Bot) DeleteWebhook()          {}
func (b *Bot) GetWebhookInfo()         {}
func (b *Bot) WebhookInfo()            {}
func (b *Bot) EditMessageText()        {}
func (b *Bot) EditMessageCaption()     {}
func (b *Bot) EditMessageReplyMarkup() {}
