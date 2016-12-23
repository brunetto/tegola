package tegola

import (
	"time"
)

type Bot struct {
	BotName string
	BotUser string
	BotToken string
	ChatId int64
	TimeZone string
}

type Updates struct {
	Ok bool `json:"ok"`
	Results []Result `json:"result"`
}

type Result struct {
	UpdateId int64 `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	MessageId int64 `json:"message_id"`
	Timestamp int64 `json:"date"`
	DateTime string
	Text string `json:"text"`
	From From `json:"from"`
	Chat Chat `json:"chat"`
}

func (m *Message) ReadTime (timezone string) error {
	var (
		tz *time.Location
		err error
	)
	tz, err = time.LoadLocation(timezone)
	if err != nil {
		return err
	}
	m.DateTime = time.Unix(m.Timestamp, 0).In(tz).String()
	return err
}

type From struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	Username string `json:"username"`
}

type Chat struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	Username string `json:"username"`
	Type string `json:"type"`
}

type outGoingMsg struct {
	ChatId int64 `json:"chat_id"`
	Text string `json:"text"`
}