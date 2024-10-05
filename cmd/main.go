package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const command = "Че сожрать"

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(command),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7624316444:AAHWLQNfzGOjzjf0p99l8WHe5nwPcHXpGZQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text != command {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Жми кнопку жэсть")
			msg.ReplyMarkup = numericKeyboard
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message: %s", err)
			}
			continue
		}

		if update.Message.Text == command {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, Jri())
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message: %s", err)
			}
			continue
		}
	}
}
