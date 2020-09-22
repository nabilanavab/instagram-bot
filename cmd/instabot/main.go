package main

import(
bot "InstaFollower/internal/app/instabot"
"log"
)

func main() {
	bot, err := bot.CreateBot()
	if err != nil {
		log.panic(err)
	}

	bot.Run()
}
