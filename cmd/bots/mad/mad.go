package mad

import (
	tg "github.com/brunetto/tegola"
)

type Lists []List

type List struct {
	Name string
	Items []Item
	Creator tg.User
	CreationDate string
}

type Item struct {
	Text string
	Status string
}

func (l *List) MarkAsDone(idx int) () {
	l.Items[idx].MarkAsDone()
}

func (i *Item) MarkAsDone() () {
	i.MarkAs("Done")
}

func (i *Item) MarkAs(status string) {
	i.Status = status
}
