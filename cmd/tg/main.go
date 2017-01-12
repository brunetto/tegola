package main

import (
	"github.com/brunetto/goutils/debug"
	tg "github.com/brunetto/tegola"
	"log"
)

var Debug = true

func main() {

	var (
		b   tg.Bot
	)

	debug.LogDebug(Debug, "Load bot from file: tegola.json")
	b = tg.NewBotFromJsonFile("tegola.json")

	b.Debug = Debug
	b.ListenRoute = "/" // "/" + b.BotToken + "/"

	err := b.Start(tg.Echo)
	if err != nil {
		log.Fatal(err)
	}
}


