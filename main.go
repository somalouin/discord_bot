package main

import (
	"discord_bot/bot"
	"discord_bot/config"
	"fmt"
)

func main() {
	fmt.Println("bot starting now")
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
