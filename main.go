package main

import (
	"packrat/backend"
	"packrat/bot"
)

func main() {
	backend.Server()
	bot.Start()
}
