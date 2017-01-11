package tegola

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/*
import (
	"net/http"
	"encoding/json"
)

func (b *Bot) ReceiveWebhookUpdates (w http.ResponseWriter, r *http.Request) {
	var (
		allowed, forbidden []Update
		err error
		b Bot
		reply Message
		lastUpdateId int64
		u  Updates
	)

	u = Updates{}
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return allowed, forbidden, err
	}
}
*/

func (b *Bot) Handler(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		reply Message
		u     Update
	)

	log.Println("init updates")

	u = Update{}
	log.Println("decode updates")

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(u, err)
	}

	messageText := u.Message.Text
	chatId := strconv.Itoa(int(u.Message.Chat.Id))
	sender := u.Message.From
	date, err := u.Message.UnixToHumanDate(b.TimeZone)

	if err != nil {
		log.Println(err)
	}

	// Echo to admin
	replyText := "Echo" + "\n====" +
		"\nSender: " + sender.Username + " - id: " + strconv.Itoa(int(sender.Id)) +
		"\nChat: " + chatId +
		"\nTimestamp " + date +
		"\nUpdate n.: " + strconv.Itoa(int(u.UpdateId)) +
		"\nMessage n.: " + strconv.Itoa(int(u.Message.MessageId)) +
		"\nMessage:\n " + messageText + "\n"

	sp := SendMessagePayload{
		ChatId: b.AdminChats[0],
		Text:   replyText,
	}

	reply, err = b.SendMessage(sp)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Echoed message sent to chat " + strconv.Itoa(int(b.AdminChats[0])) + " is: ")
	fmt.Println(reply.Text)

}

func (b *Bot) SetWebhook()     {}
func (b *Bot) DeleteWebhook()  {}
func (b *Bot) GetWebhookInfo() {}
func (b *Bot) WebhookInfo()    {}
