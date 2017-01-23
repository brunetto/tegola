package tegola

import (
	"github.com/brunetto/goutils/conf"
	"log"
	"net/http"
	"time"
	"golang.org/x/tools/go/gcimporter15/testdata"
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
	b.Init()
	return b
}

func NewBotsFromJsonFileList(fileNames []string) []Bot {
	var (
		bot = Bot{}
		bots   = []Bot{}
	)

	for _, fileName := range fileNames {
		bot = NewBotFromJsonFile(fileName)
		bot.Init()
		bots = append(bots, bot)
	}

	return bots
}

func (b *Bot)Init() {
	b.Client = &http.Client{Timeout: time.Second * 10}
	b.UpdatesChan = make(chan Update, b.UpdatesChanSize)
	b.ListenRoute = "/" + b.BotToken + "/"
	b.ListenPort = "8443"
}
