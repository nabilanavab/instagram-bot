package main

import(
bot "InstaFollower/internal/app/instabot"
"log"
)

func main() {
	bot, err := telegram.CreateBot(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.print(err)
		return
	}

	bot.Run()
}
