package main

import (
	"InstaFollower/internal/instabot/bot.go"
	"InstaFollower/internal/instabot/utils"
)

func main() {
	config := utils.CreateConfig()

	db, err := db.CreateConnection(config.PostgresURI)
	if err != nil {
		Print(err)
		return
	}
	defer db.Disconnect()

	bot, err := bot.CreateBot(config, db)
	if err != nil {
		Print(err)
		return
	}

	bot.Run()
}
