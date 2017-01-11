package main

import (
	"github.com/brunetto/goutils/debug"
	tg "github.com/brunetto/tegola"
	"log"
	"net/http"
)

var Debug = true

func main() {

	var (
		b   tg.Bot
		err error
	)

	debug.LogDebug(Debug, "Load bot from file: tegola.json")
	b = tg.NewBotFromJsonFile("tegola.json")

	b.Debug = Debug

	debug.LogDebug(Debug, "Setting up handler")

	http.HandleFunc("/" /*+b.BotToken+"/"*/, b.Handler)

	debug.LogDebug(Debug, "Listening for updates from webhook")

	err = http.ListenAndServe("127.0.0.1:8443", nil)
	if err != nil {
		log.Fatal(err)
	}
}
