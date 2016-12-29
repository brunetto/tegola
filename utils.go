package tegola

import "time"

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
