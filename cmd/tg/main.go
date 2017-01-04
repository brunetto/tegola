package main

import (
	tg "github.com/brunetto/tegola"
	"log"
	"fmt"
	"strconv"
	"github.com/brunetto/goutils/debug"
	"time"
)

var Debug = false

func main () {
	var (
		allowed, forbidden []tg.Update
		err error
		b tg.Bot
		reply tg.Message
		lastUpdateId int64
	)


	debug.LogDebug(Debug, "Load bot from file: tegola.json" )
	b = tg.NewBotFromJsonFile("tegola.json")

	b.Debug = Debug

	debug.LogDebug(Debug, "Getting updates")


	for {

		gp := tg.GetUpdatesPayload{Offset: lastUpdateId + 1}

		allowed, forbidden , err = b.GetUpdates(gp)
		if err != nil {
			log.Fatal(err)
		}

		//debug.LogDebug(Debug, "Updates checked")

		for _, u := range allowed {
			lastUpdateId = u.UpdateId
			messageText := u.Message.Text
			chatId := strconv.Itoa(int(u.Message.Chat.Id))
			sender := u.Message.From
			date, err := u.Message.UnixToHumanDate(b.TimeZone)

			if err != nil {
				log.Fatal(err)
			}

			replyText := "Echo" + "\n====" +
					"\nSender: " + sender.Username + " - id: " + strconv.Itoa(int(sender.Id)) +
					"\nChat: " + chatId +
					"\nTimestamp " + date +
					"\nUpdate n.: " + strconv.Itoa(int(u.UpdateId)) +
					"\nMessage n.: " + strconv.Itoa(int(u.Message.MessageId)) +
					"\nMessage:\n " + messageText + "\n"

			sp := tg.SendMessagePayload{
				ChatId: u.Message.Chat.Id,
				Text: replyText,
			}

			reply, err = b.SendMessage(sp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Echoed message sent to chat is " + strconv.Itoa(int(u.Message.Chat.Id)) + " is: ")
			fmt.Println(reply.Text)

		}

		for _, u := range forbidden {
			lastUpdateId = u.UpdateId
			messageText := u.Message.Text
			chatId := strconv.Itoa(int(u.Message.Chat.Id))
			sender := u.Message.From
			date, err := u.Message.UnixToHumanDate(b.TimeZone)

			if err != nil {
				log.Fatal(err)
			}

			// Echo to admin
			replyText := "Echo" + "\n====" +
				"\nSender: " + sender.Username + " - id: " + strconv.Itoa(int(sender.Id)) +
				"\nChat: " + chatId +
				"\nTimestamp " + date +
				"\nUpdate n.: " + strconv.Itoa(int(u.UpdateId)) +
				"\nMessage n.: " + strconv.Itoa(int(u.Message.MessageId)) +
				"\nMessage:\n " + messageText + "\n"

			sp := tg.SendMessagePayload{
				ChatId: b.AdminChats[0],
				Text: replyText,
			}

			reply, err = b.SendMessage(sp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Echoed message sent to chat " + strconv.Itoa(int(b.AdminChats[0])) + " is: ")
			fmt.Println(reply.Text)


			// Notify the bot will ask for permissison
			sp2 := tg.SendMessagePayload{
				ChatId: u.Message.Chat.Id,
				Text: "I can't talk to strangers, I'll ask for permission to my Admin",
			}

			reply, err = b.SendMessage(sp2)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Echoed message sent to chat " + strconv.Itoa(int(u.Message.Chat.Id)) + " is: ")
			fmt.Println(reply.Text)

		}

		time.Sleep(1*time.Second)

	}

}
