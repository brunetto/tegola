package tegola

import (
	"encoding/json"
	"net/http"
)

// @todo: find out if it is bettere to recycle the http client or not

func (b *Bot) SimplerGetUpdates() ([]Update, []Update, error) {
	var (
		resp      *http.Response
		err       error
		u         Updates
		allowed   = []Update{}
		forbidden = []Update{}
	)
	resp, err = b.Get("getUpdates")
	defer resp.Body.Close()
	if err != nil {
		return allowed, forbidden, err
	}

	u = Updates{}
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return allowed, forbidden, err
	}

	allowed, forbidden = b.FilterAllowedUpdates(u)

	return allowed, forbidden, err
}

func (b *Bot) GetUpdates(pReq GetUpdatesPayload) ([]Update, []Update, error) {
	var (
		resp      *http.Response
		err       error
		u         Updates
		allowed   = []Update{}
		forbidden = []Update{}
		payload []byte
	)

	payload, err = json.Marshal(pReq)
	if err != nil {
		return allowed, forbidden, err
	}

	resp, err = b.Post("getUpdates", payload)
	defer resp.Body.Close()
	if err != nil {
		return allowed, forbidden, err
	}

	u = Updates{}
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return allowed, forbidden, err
	}

	allowed, forbidden = b.FilterAllowedUpdates(u)

	return allowed, forbidden, err
}

func (b *Bot) FilterAllowedUpdates(u Updates) ([]Update, []Update) {
	var (
		allowed   = []Update{}
		forbidden = []Update{}
	)
	for _, Update := range u.Updates {
		if b.AllowedMessage(Update.Message) {
			allowed = append(allowed, Update)
		} else {
			forbidden = append(forbidden, Update)
		}
	}
	return allowed, forbidden
}

func (b *Bot) SendSimpleMessage(msgText string) (Message, error) {
	var (
		msgReq = SendMessagePayload{ChatId: b.ChatId, Text: msgText}
		m      Message
		err    error
	)
	m, err = b.SendMessage(msgReq)
	return m, err
}

func (b *Bot) SendMessage(msgReq SendMessagePayload) (Message, error) {
	var (
		resp *http.Response
		err  error
		msg  []byte
		m    = Message{}
	)

	msg, err = json.Marshal(msgReq)
	if err != nil {
		return m, err
	}

	resp, err = b.Post("sendMessage", msg)
	defer resp.Body.Close()
	if err != nil {
		return m, err
	}

	// @todo: fix
	err = json.NewDecoder(resp.Body).Decode(&m)
	return m, err
}

func (b *Bot) SetWebhook() {}

func (b *Bot) GetMe()                  {}
func (b *Bot) DeleteWebhook()          {}
func (b *Bot) GetWebhookInfo()         {}
func (b *Bot) WebhookInfo()            {}
func (b *Bot) EditMessageText()        {}
func (b *Bot) EditMessageCaption()     {}
func (b *Bot) EditMessageReplyMarkup() {}

func (b *Bot) ForwardMessage()        {}
func (b *Bot) SendPhoto()             {}
func (b *Bot) SendAudio()             {}
func (b *Bot) SendDocument()          {}
func (b *Bot) SendSticker()           {}
func (b *Bot) SendVideo()             {}
func (b *Bot) SendVoice()             {}
func (b *Bot) SendLocation()          {}
func (b *Bot) SendVenue()             {}
func (b *Bot) SendContact()           {}
func (b *Bot) SendChatAction()        {}
func (b *Bot) GetUserProfilePhotos()  {}
func (b *Bot) GetFile()               {}
func (b *Bot) KickChatMember()        {}
func (b *Bot) LeaveChat()             {}
func (b *Bot) UnbanChatMember()       {}
func (b *Bot) GetChat()               {}
func (b *Bot) GetChatadministrators() {}
func (b *Bot) GetChatmemberscount()   {}
func (b *Bot) GetChatmember()         {}
func (b *Bot) AnswerCallBackquery()   {}
