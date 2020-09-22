package main

import bot "InstaFollower/internal/app/instabot"

func main() {
	bot, err := bot.CreateBot()
	if err != nil {
		log.panic(err)
	}

	bot.Run()
}
