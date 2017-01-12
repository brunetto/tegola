package main

import (
	"github.com/brunetto/goutils/debug"
	tg "github.com/brunetto/tegola"
	"log"
	"net/http"
	"sync"
)

var Debug = true

func main() {

	var (
		b   tg.Bot
		err error
		wg sync.WaitGroup
	)

	debug.LogDebug(Debug, "Load bot from file: tegola.json")
	b = tg.NewBotFromJsonFile("tegola.json")

	b.Debug = Debug

	debug.LogDebug(Debug, "Setting up handler")

	http.HandleFunc("/" /*+b.BotToken+"/"*/, b.WebhooksUpdatesHandler)

	debug.LogDebug(Debug, "Listening for updates from webhook")

	wg.Add(1)
	go func(uch chan tg.Update){
		defer wg.Done()
		for u := range uch {
			b.Echo(u)
		}
	}(b.UpdatesChan)

	err = http.ListenAndServe("127.0.0.1:8443", nil)
	close(b.UpdatesChan)
	wg.Wait()
	if err != nil {
		log.Fatal(err)
	}

}
