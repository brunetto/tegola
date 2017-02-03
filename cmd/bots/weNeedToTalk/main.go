package main

import (
	tg "github.com/brunetto/tegola"
	"gitlab.com/brunetto/sdbutils"
	"log"
	"strconv"
	"sync"
	"time"
)

func main () {
	var (
		b   tg.Bot
		cr ControlRoom
		conf        sdbutils.ConnInfo
		err error
		wg  sync.WaitGroup
	)
	b = tg.NewBotFromJsonFile("wntt.json")

	cr = ControlRoom{}
	cr.db = sdbutils.NewDB()
	cr.b = &b
	cr.c = make(chan string, 10)
	conf, err = sdbutils.ReadCnf("service-users.cnf")
	if err != nil {
		log.Println(err)
	}
	cr.db.SetConnectionInfo(conf)
	err = cr.db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go cr.Notifier(&wg)
	wg.Add(1)
	go cr.CheckTimesQueue(&wg)
	//close(cr.c)
	wg.Wait()
}

type ControlRoom struct {
	db sdbutils.DB
	b *tg.Bot
	c chan string
}

func (cr *ControlRoom) CheckTimesQueue (wg *sync.WaitGroup) () {
	defer wg.Done()
	var (
		queryString string
		message string
		result      sdbutils.Records
		err error
		threshold int64 = 20
		n int64
	)
	queryString = ``
	for {
		log.Println("Starting query")
		result, err = cr.db.GetDataFromGenericTable(queryString)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Query done")
		message = "Queue lenght for cron 'Times' is: " + result[0].Fields["lavorazioni_senza_times"].Value
		n, err = strconv.ParseInt(result[0].Fields["lavorazioni_senza_times"].Value, 10,64)
		if err != nil {
			log.Fatal(err)
		}
		if n < threshold {
			message += "\nNo action required.\n"
		} else {
			message += "\nBetter check the situation.\nCheck out https://confluence.pixartprinting.com/display/IFL."
		}
		message += "\nLet's check again in one minute.\n"
		log.Println("sending ", message)
		cr.c <- message

		log.Println("Sleeping for one minute")
		time.Sleep(1*time.Minute)
		log.Println("Resume")
	}
}

func (cr *ControlRoom) Notifier (wg *sync.WaitGroup) () {
	defer wg.Done()
	for message := range cr.c {
		log.Println(message)
		_, err := cr.b.SendSimpleMessage(0, message)
		if err != nil {
			log.Fatal(err)
		}
	}
}

