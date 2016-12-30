package main

import (
	tg "github.com/brunetto/tegola"
	"log"
	"fmt"
)

func main () {
	var (
		allowed, forbidden []tg.Update
		err error
		b tg.Bot
		//reply tg.Message
	)
	b = tg.NewBotFromJsonFile("tegola.json")
	allowed, forbidden, err = b.SimplerGetUpdates()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Allowed:\n%+v\n", allowed)
	fmt.Printf("Forbidden:\n%+v\n", forbidden)

	/*reply*/_, err = b.SendSimpleMessage("Test")
	if err != nil {
		log.Fatal(err)
	}


}
