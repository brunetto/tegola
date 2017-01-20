package expenseTracker

import (
	tg "github.com/brunetto/tegola"
)

type Item struct {
	Currency string
	Amount float64
	DateTime string
	Description string
	Tags []string
	Type string
}

type account struct {
	Name string
	Owner tg.User
	Items []Item
}
