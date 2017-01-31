package tegola

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/brunetto/goutils/debug"
)

func (b *Bot) GenericMethod(method string, payload []byte) ([]byte, error) {
	var (
		resp     *http.Response
		err      error
		response []byte
	)

	resp, err = b.Post(method, payload)
	defer resp.Body.Close()
	if err != nil {
		return response, err
	}

	// Read the response into a byte array
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	return response, err
}

func (b *Bot) SimplerGetUpdates() ([]Update, error) {
	var (
		resp    *http.Response
		err     error
		u       Updates
		updates = []Update{}
	)
	debug.LogDebug(b.Debug, "Get request")
	resp, err = b.Get("getUpdates")
	defer resp.Body.Close()
	if err != nil {
		return updates, errors.New("Failed getUpdates get requests: " + err.Error())
	}

	u = Updates{}

	debug.LogDebug(b.Debug, "Decode response")
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return updates, err
	}

	debug.LogDebug(b.Debug, "Filter updates")
	//allowed, forbidden = b.FilterAllowedUpdates(u)

	return u.Updates, err
}

func (b *Bot) GetUpdates(pReq GetUpdatesPayload) ([]Update, error) {
	var (
		resp    *http.Response
		err     error
		u       Updates
		updates = []Update{}
		payload []byte
	)

	payload, err = json.Marshal(pReq)
	if err != nil {
		return updates, err
	}

	// Start option 1
	resp, err = b.Post("getUpdates", payload)
	defer resp.Body.Close()
	if err != nil {
		return updates, err
	}

	u = Updates{}
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return updates, err
	}
	// End option 1

	// Start option 2
	/*
		var respB []byte
		respB, err = b.GenericMethod("getUpdates", payload)
		if err != nil {
			return allowed, forbidden, err
		}
		u = Updates{}
		err = json.Unmarshal(respB, &u)
		if err != nil {
			return allowed, forbidden, err
		}
	*/
	// End option 2

	// Now GetUpdates returns all the updates, they are going to be filtered after
	//allowed, forbidden = b.FilterAllowedUpdates(u)

	return u.Updates, err
}

func (b *Bot) FilterAllowedUpdates(u []Update) ([]Update, []Update) {
	var (
		allowed   = []Update{}
		forbidden = []Update{}
	)
	for _, Update := range u {
		if b.AllowedMessage(Update.Message) {
			allowed = append(allowed, Update)
		} else {
			forbidden = append(forbidden, Update)
		}
	}
	return allowed, forbidden
}

func (b *Bot) SendSimpleMessage(chatId int64, msgText string) (Message, error) {
	var (
		msgReq = SendMessagePayload{ChatId: chatId, Text: msgText}
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
		m    = SendMessageConfirm{}
	)

	msg, err = json.Marshal(msgReq)
	if err != nil {
		return Message{}, err
	}

	resp, err = b.Post("sendMessage", msg)
	defer resp.Body.Close()
	if err != nil {
		return Message{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&m)
	return m.Message, err
}

func (b *Bot) GetMe()                  {}
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
