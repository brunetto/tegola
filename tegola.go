package tegola

import (
	"net/http"
	"time"
	"encoding/json"
	"bytes"
)

func (b *Bot) GetUpdates () ([]Result, []Result, error) {
	var (
		client = &http.Client{Timeout: time.Second * 10}
		resp *http.Response
		err error
		u Updates
		allowed = []Result{}
		forbidden = []Result{}
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

	for _, result := range u.Results {
		if b.AllowedMessage(result.Message) {
			allowed = append(allowed, result)
		} else {
			forbidden = append(forbidden, result)
		}
	}
	return allowed, forbidden, err
}

func (b *Bot) SendSimpleMsg(msgText string) (Message, error) {
	var (
		client = &http.Client{Timeout: time.Second * 10}
		resp *http.Response
		err error
		m Message
		msg []byte
	)

	m = Message{}

	msg, err = json.Marshal(OutGoingMsg{b.ChatId, msgText})
	if err != nil {
		return m, err
	}


	resp, err = client.Post("https://api.telegram.org/bot" + b.BotToken + "/sendMessage", "application/json", bytes.NewBuffer(msg))
	if err != nil {
		return m, err
	}

	err = json.NewDecoder(resp.Body).Decode(&m)
	return m, err
}

