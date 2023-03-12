package main

import (
	"fmt"

	"github.com/gabrielm3/bot-discord/bot"
	"github.com/gabrielm3/bot-discord/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return

}
