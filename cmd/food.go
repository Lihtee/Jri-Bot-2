package main

import (
	wr "github.com/mroth/weightedrand"
	"log"
)

type Food struct {
	Name   string
	Weight int
}

var jri []*Food = []*Food{
	newFood("Сожри пицц", 2),               // 2 в сесяц
	newFood("Сожри роллс", 1),              // 1 в сесяц
	newFood("Сожри курицу из KFC", 8),      // 2 раза в неделю
	newFood("Сожри чикан", 0),              // 0 раз (только с сырной коллеццией можна жрать это, в последний сас была зимой)
	newFood("Сожри арбыс или дыню", 1),     // 1 раз в сесяц
	newFood("Сожри бургерс", 8),            // 1 раз в неделю
	newFood("Сожри шашлык", 4),             // 1 раз в неделю
	newFood("Сожри шаурму", 4),             // 1 раз в неделю
	newFood("Сожри хинкали с хачапури", 1), // 1 раз в месяц
	newFood("Сожри сэсвич", 2),             // 2 раза в сесяц (осталось найти где их жрать)
	newFood("Сожри стейк", 1),              // 1 раз в сесяц
	newFood("Сожри блин", 2),               // 2 раза в сесяц
	newFood("Сожри пак гадюки", 1),         // 1 раз в сесяц
	newFood("Сожри ночной снэк", 8),        // 2 раза в неделю
	newFood("Сожри рамен", 4),              // 1 раз в неделю
	newFood("Сожри WOK", 4),                // 1 раз в неделю
	newFood("Сожри рыс с яйцом", 4),        // 1 раз в неделю
	newFood("Сожри сало", 2),               // 2 раза в сесяц
	newFood("Сожри CUMнам", 1),             // 1 раз в сесяц
	newFood("Сожри чабуреки", 2),           // 2 раза в сесяц
	newFood("Сожри пироженку", 1),          // 1 раз в сесяц
	newFood("Сожри колбаски", 2),           // 2 раза в сесяц
	newFood("Сожри понтовые вафли", 1),     // 1 раз в сесяц
	newFood("Сожри карбонару", 2),          // 2 раза в сесяц
	newFood("Сожри торт", 1),               // 1 раз в сесяц
}

var chooser = newChooser()

func Zri() string {
	return chooser.Pick().(string)
}

func (food *Food) toChoice() wr.Choice {
	return wr.Choice{Item: food.Name, Weight: uint(food.Weight)}
}

func newChooser() *wr.Chooser {
	choices := []wr.Choice{}
	for _, food := range jri {
		choices = append(choices, food.toChoice())
	}
	chooser, err := wr.NewChooser(choices...)
	if err != nil {
		log.Panicf("failed to init food: %s", err)
	}

	return chooser
}

func newFood(name string, weight int) *Food {
	return &Food{name, weight}
}
