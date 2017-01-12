package tegola

import (
	"encoding/json"
	"log"
	"net/http"
)

func (b *Bot) WebhooksUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		u   Update
	)

	log.Println("Init updates")

	u = Update{}
	log.Println("Decode updates")

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(u, err)
	}

	// Send update down to the bot updates channel
	b.UpdatesChan <- u
}

func (b *Bot) SetWebhook()     {}
func (b *Bot) DeleteWebhook()  {}
func (b *Bot) GetWebhookInfo() {}
func (b *Bot) WebhookInfo()    {}
