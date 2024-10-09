package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

const command = "Че сожрать"

var (
	jriButton = tele.ReplyButton{Text: command}
	keyboard  = &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{jriButton},
		},
	}
)

// Unique is best to be english letters only, otherwise the regex inside router breaks
var (
	spiButton     = tele.InlineButton{Text: "spi", Unique: "spi"}
	neSpiButton   = tele.InlineButton{Text: "ne spi", Unique: "neSpi"}
	inlineSpiMenu = &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{spiButton},
		},
	}
	inlineNeSpiMenu = &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{
			{neSpiButton},
		},
	}
)

func main() {
	pref := tele.Settings{
		//Token:  os.Getenv("TOKEN"),
		Token:   "7624316444:AAHWLQNfzGOjzjf0p99l8WHe5nwPcHXpGZQ",
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: false,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(command, func(c tele.Context) error {
		return c.Send(Jri(), inlineSpiMenu)
	})

	b.Handle("\fspi", func(c tele.Context) error {
		fmt.Println("spi button_____")
		_, err := b.Edit(c.Message(), inlineNeSpiMenu)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		return err
	})

	b.Handle(&neSpiButton, func(c tele.Context) error {
		fmt.Println("ne spi button_____")
		_, err := b.Edit(c.Message(), inlineSpiMenu)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		return err
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("Жми кнопку жэсть", keyboard)
	})

	b.Start()
}
