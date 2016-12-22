package tegola

import (
	"net/http"
	"os"
	"log"
	"io"
)

type Bot struct {
	BotName string
	BotUser string
	BotToken string
}

func (b *Bot) GetUpdates () {
	response, err := http.Get("https://api.telegram.org/bot" + b.BotToken + "/getUpdates")
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}

