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

var defaultMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("spi", "spi"),
	),
)

var updatedMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ne spi", "ne spi"),
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
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			respondToMessage(update.Message, bot)
		}

		if update.CallbackQuery != nil {
			// Some logging
			//j, err := json.Marshal(update)
			//if err != nil {
			//	log.Printf("Error on marshaling update: %s", err)
			//} else {
			//	log.Printf("Update is: %s", string(j))
			//}
			handleButtonClick(update.CallbackQuery, bot)
		}
	}
}

func handleButtonClick(query *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	if query.Data == "spi" {
		message := query.Message
		update := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, message.Text, updatedMenuMarkup)
		_, err := bot.Send(update)
		if err != nil {
			log.Printf("Error handling %s callback: %s", query.Data, err)
		}
		return
	}

	if query.Data == "ne spi" {
		message := query.Message
		update := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, message.Text, defaultMenuMarkup)
		_, err := bot.Send(update)
		if err != nil {
			log.Printf("Error handling %s callback: %s", query.Data, err)
		}
		return
	}
}

func respondToMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	if message.Text != command {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Жми кнопку жэсть")
		msg.ReplyMarkup = numericKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %s", err)
		}
		return
	}

	if message.Text == command {
		msg := tgbotapi.NewMessage(message.Chat.ID, Jri())
		msg.ReplyMarkup = defaultMenuMarkup
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %s", err)
		}
		return
	}
}
