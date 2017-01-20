package tegola

import (
	"github.com/brunetto/goutils/conf"
	"log"
	"net/http"
	"time"
)

func NewBotFromJsonFile(fileName string) Bot {
	var (
		err error
		b   = Bot{}
	)
	err = conf.LoadJsonConf(fileName, &b)
	if err != nil {
		log.Fatal("Error reading JSON config file: ", err)
	}
	b = InitBot(b)
	return b
}

func InitBot(b Bot) Bot {
	b.Client = &http.Client{Timeout: time.Second * 10}
	b.UpdatesChan = make(chan Update, b.UpdatesChanSize)
	b.ListenRoute = "/"
	//b.CommandRegString = `^\/(?P<command>[^@\s]+)@?(?:(?P<bot>\S+)|)\s?(?P<args>[\s\S]*)$`
	return b
}
