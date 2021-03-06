package mad

import (
	tg "github.com/brunetto/tegola"
	"log"
	"strings"
)

type Lists map[string]List

type ItemsList struct {
	Name string
	Items []Item
	Creator tg.User
	Chat tg.Chat
	CreationDate string
}

type Item struct {
	Text string
	Status string
}

func (l *ItemsList) MarkItemAsDone(i int) () {
	l.Items[i].MarkAsDone()
}

func (i *Item) MarkAsDone() () {
	i.MarkAs("Done")
}

func (i *Item) MarkAs(status string) {
	i.Status = status
}

func (l *ItemsList) DeleteItem(i int) {
	l.Items = append(l.Items[:i], l.Items[i+1:]...)
}

func LoadCommands () (*tg.CmdManager, error) {
	var (
		c = tg.NewCmdManager()
		err error
	)

	err = c.AddRoute("add", Add, true)
	err = c.AddRoute("mad", Mad, true)
	err = c.AddRoute("del", Del, true)
	err = c.AddRoute("list", List, true)
	err = c.AddRoute("clean", Clean, true)
	//err = c.AddRoute("newList", NewList, true)
	//err = c.AddRoute("delList", DelList, true)

	return c, err
}

func Add (b *tg.Bot, c tg.CmdData, u tg.Update) error {
	var (
		err error
	)

	list := strings.Split(u.Message.Text, ",")

	for _, item := range list {
		// add item to list
		log.Println(item)
	}

	// save list

	replyText := "Added: " + strings.Join(list, ", ")

	sp := tg.SendMessagePayload{
		ChatId: u.Message.Chat.Id,
		Text:   replyText,
	}

	_, err = b.SendMessage(sp)
	if err != nil {
		log.Println(err)
	}
	return err
}

func Mad (b *tg.Bot, c tg.CmdData, u tg.Update) error {
	var (
		err error
	)
	return err
}
func Del (b *tg.Bot, c tg.CmdData, u tg.Update) error {
	var (
		err error
	)
	return err
}
func List (b *tg.Bot, c tg.CmdData, u tg.Update) error {
	var (
		err error
	)
	return err
}
func Clean (b *tg.Bot, c tg.CmdData, u tg.Update) error {
	var (
		err error
	)
	return err
}
//func NewList (b *tg.Bot, c tg.CmdData, u tg.Update) error {
//	var (
//		err error
//	)
//	return err
//}
//func DelList (b *tg.Bot, c tg.CmdData, u tg.Update) error {
//	var (
//		err error
//	)
//	return err
//}




