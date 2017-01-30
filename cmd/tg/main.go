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
		c *tg.CmdManager
		err error
	)

	debug.LogDebug(Debug, "Load bot from file: tegola.json")
	b = tg.NewBotFromJsonFile("tegola.json")

	b.Debug = Debug
	b.UpdateMode = "getUpdates" // To interact during local tests

	c = tg.NewCmdManager()

	//err = c.ModifyRoute("default", tg.EchoHandler)
	//if err != nil {
	//	log.Fatal(err)
	//}

	err = c.AddRoute("start", tg.SayHi, true)
	if err != nil {
		log.Fatal(err)
	}

	err = c.AddRoute("help", tg.Help, true)
	if err != nil {
		log.Fatal(err)
	}

	err = b.Start(c.CmdRouter)
	if err != nil {
		log.Fatal(err)
	}
}
