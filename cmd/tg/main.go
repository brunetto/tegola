package main

import (
	tg "github.com/brunetto/tegola"
)

func main () {
	b := tg.LoadBot("tegola.json")
	b.GetUpdates()
}
