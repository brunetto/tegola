package tegola

import (
	"time"
	"github.com/brunetto/figa"
)

type Bot struct {
	BotName string
	BotUser string
	BotToken string
	ChatId int64
	TimeZone string
	Admin []User
	AllowedUsers []User
	AllowedChats []int64
	FAppConf figa.FApp
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
	Location Location `json:"location"`
	From User `json:"from"`
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

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
}

type Chat struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"` // Optional. First name of the other party in a private chat
	Username string `json:"username"` // Optional. Username, for private chats, supergroups and channels if available
	Type string `json:"type"` // can be either “private”, “group”, “supergroup” or “channel”
	Title string `json:"title"` // Optional. Title, for supergroups, channels and group chats
	LastName string `json:"last_name"` //		Optional. Last name of the other party in a private chat
	AllMembersAreAdministrators bool `json:"all_members_are_administrators"`//		Optional. True if a group has ‘All Members Are Admins’ enabled.
}
//https://core.telegram.org/bots/api#available-types
type OutGoingMsg struct {
	ChatId int64 `json:"chat_id"`
	Text string `json:"text"`
}

type Location struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (b *Bot) IsAllowedChat ( chatId int64 ) bool {
	for _, id := range b.AllowedChats {
		if chatId == id {
			return true
		}
	}
	return false
}

func (b *Bot) IsAllowedUser ( from User ) bool {
	for _, user := range b.AllowedUsers {
		if from.Id == user.Id && from.Username == user.Username {
			return true
		}
	}
	return false
}

func (b *Bot) AllowedMessage(m Message) bool {
	allowed := b.IsAllowedChat(m.Chat.Id) && b.IsAllowedUser(m.From)
	return allowed
}
