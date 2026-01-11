package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dreyspi/jribot2/jri"

	tele "gopkg.in/telebot.v4"
)

const jriCommand = "Че сожрать"
const star = "⭐"

var (
	jriButton   = tele.ReplyButton{Text: jriCommand}
	replyLayout = &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{jriButton},
		},
	}
)

func main() {
	pref := tele.Settings{
		Token:   os.Getenv("TOKEN"),
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: false,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = b.SetCommands([]tele.Command{
		{Text: "/start", Description: "Хочу жрать"},
		{Text: "/packs", Description: "Покажи еду"},
	})

	if err != nil {
		log.Printf("could not set commands: %s", err)
		// Works ok without commands
	}

	b.Handle(jriCommand, func(c tele.Context) error {
		food, err := jri.Jri(c.Sender().ID)
		if err != nil {
			return fmt.Errorf("failed to get food: %w", err)
		}

		return c.Send(food, replyLayout)
	})

	b.Handle("/start", func(c tele.Context) error {
		msg := fmt.Sprintf("Жэээээсть, %s опять жрать хочет", c.Sender().Username)

		return c.Send(msg, replyLayout)
	})

	b.Handle("/packs", func(c tele.Context) error {
		msg := fmt.Sprintf("Жэээээсть, %s смотри че есть и жми", c.Sender().Username)
		selectedPresetId, err := jri.Eda(c.Sender().ID)
		if err != nil {
			return fmt.Errorf("failed to get selectedPresetId: %w", err)
		}

		return c.Send(msg, presetMenuLayout(selectedPresetId))
	})

	b.Handle("\fpack", func(c tele.Context) error {
		packId := c.Callback().Data
		err := jri.SetEda(c.Sender().ID, packId)
		if err != nil {
			return err
		}

		_, err = b.Edit(c.Message(), presetMenuLayout(packId))
		if err != nil {
			// Not good but packId is saved for user so send food anyway
			fmt.Printf("falied to update inline menu: %s", err)
			// Send something to unlock the button
			_ = c.Respond(&tele.CallbackResponse{})
		}

		food, err := jri.Jri(c.Sender().ID)
		if err != nil {
			return err
		}

		return c.Send(food, replyLayout)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("Жми кнопку жэсть", replyLayout)
	})

	fmt.Println("Starting")
	b.Start()
}

func presetMenuLayout(packId string) *tele.ReplyMarkup {

	// Unique is best to be english letters only, otherwise the regex inside router breaks
	buttons := [][]tele.InlineButton{
		{
			tele.InlineButton{Text: getStar(packId, jri.BasedPresetId) + " Базированный пак", Unique: "pack", Data: jri.BasedPresetId},
			tele.InlineButton{Text: "Че тут", Unique: "che", Data: jri.BasedPresetId},
		},
		{
			tele.InlineButton{Text: getStar(packId, jri.ThaiPresetId) + "Тайский пак", Unique: "pack", Data: jri.ThaiPresetId},
			tele.InlineButton{Text: "Че тут", Unique: "che", Data: jri.ThaiPresetId},
		},
		{
			tele.InlineButton{Text: getStar(packId, jri.GeorgianPresetId) + "Грузинский пак", Unique: "pack", Data: jri.GeorgianPresetId},
			tele.InlineButton{Text: "Че тут", Unique: "che", Data: jri.GeorgianPresetId},
		},
	}

	return &tele.ReplyMarkup{InlineKeyboard: buttons}
}

func getStar(selectedPackId string, menuPackId string) string {
	if selectedPackId == menuPackId {
		return star
	}

	return ""
}
