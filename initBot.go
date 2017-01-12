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
		c   = Bot{}
	)
	err = conf.LoadJsonConf(fileName, &c)
	if err != nil {
		log.Fatal("Error reading JSON config file: ", err)
	}
	c = InitBot(c)
	return c
}

func InitBot(c Bot) Bot {
	c.Client = &http.Client{Timeout: time.Second * 10}
	c.UpdatesChan = make(chan Update, c.UpdatesChanSize)
	return c
}
