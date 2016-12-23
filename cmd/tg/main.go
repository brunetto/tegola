package main

import (
	tg "github.com/brunetto/tegola"
	"log"
	"fmt"
)

func main () {
	var (
		updates tg.Updates
		err error
		b tg.Bot
		reply tg.Message
	)
	b = tg.LoadBot("tegola.json")
	updates, err = b.GetUpdates()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", updates)

	reply, err = b.SendSimpleMsg("Test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", reply)

}
