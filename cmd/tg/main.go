package main

import (
	"github.com/brunetto/goutils/debug"
	tg "github.com/brunetto/tegola"
	"log"
)

var Debug = false

func main() {

	var (
		b   tg.Bot
		c *tg.CmdManager
		err error
	)

	debug.LogDebug(Debug, "Load bot from file: tegola.json")
	b = tg.NewBotFromJsonFile("tegola.json")

	b.Debug = Debug
	b.ListenRoute = "/" // "/" + b.BotToken + "/"

	//err = b.Start(tg.Echo)
	log.Println("create new cmd mnager")
	c = tg.NewCmdManager()
	log.Println("add default route")
	err = c.ModifyRoute("default", tg.EchoHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = c.AddRoute("start", SayHi, true)
	if err != nil {
		log.Fatal(err)
	}

	err = c.AddRoute("help", Help, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start")

	err = b.Start(c.CmdRouter)
	if err != nil {
		log.Fatal(err)
	}
}


func SayHi(b *tg.Bot, c tg.CmdData, u tg.Update) error {
	var (
		err error
	)

	sender := u.Message.From

	if err != nil {
		log.Println(err)
	}


	replyText := "Hi " + sender.Username + " !\n"

	sp := tg.SendMessagePayload{
		ChatId: u.Message.Chat.Id,
		Text:   replyText,
	}

	_, err = b.SendMessage(sp)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func Help(b *tg.Bot, c tg.CmdData, u tg.Update) error {
	var (
		err error
	)

	if err != nil {
		log.Println(err)
	}


	replyText := "Usage here\n"

	sp := tg.SendMessagePayload{
		ChatId: u.Message.Chat.Id,
		Text:   replyText,
	}

	_, err = b.SendMessage(sp)
	if err != nil {
		log.Println(err)
	}
	return nil
}


