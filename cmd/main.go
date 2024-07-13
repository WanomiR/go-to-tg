package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0) // set initial offset
	updateConfig.Timeout = 30             // wait up to 30 seconds on each request
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		// Telegram can send many types of updates. For now, we only want to look at messages
		if update.Message == nil {
			continue
		}

		// now that we've gotten a message, construct a reply
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err = bot.Send(msg); err != nil {
			log.Println(err)
			continue
		}
	}

}
