package main

import (
	fr "github.com/dreyspi/jribot2/cmd/frequency"
	wr "github.com/mroth/weightedrand"
	"log"
)

type Food struct {
	Name   string
	Weight int
}

var jri []*Food = []*Food{
	newFood("Сожри пицц", fr.Month*2),
	newFood("Сожри роллс", fr.Month*1),
	newFood("Сожри курицу из KFC", fr.Week*2),
	newFood("Сожри чикан", 0),
	newFood("Сожри арбыс или дыню", fr.Month*1),
	newFood("Сожри бургерс", fr.Week*1),
	newFood("Сожри шашлык", fr.Week*1),
	newFood("Сожри шаурму", fr.Week*1),
	newFood("Сожри хинкали с хачапури", fr.Month*1),
	newFood("Сожри сэсвич", fr.Month*2),
	newFood("Сожри стейк", fr.Month*1),
	newFood("Сожри блин", fr.Month*2),
	newFood("Сожри пак гадюки", fr.Month*1),
	newFood("Сожри ночной снэк", fr.Week*2),
	newFood("Сожри рамен", fr.Week*1),
	newFood("Сожри WOK", fr.Week*1),
	newFood("Сожри рыс с яйцом", fr.Week*1),
	newFood("Сожри сало", fr.Month*2),
	newFood("Сожри CUMнам", fr.Month*1),
	newFood("Сожри чабуреки", fr.Month*2),
	newFood("Сожри пироженку", fr.Month*1),
	newFood("Сожри колбаски", fr.Month*2),
	newFood("Сожри понтовые вафли", fr.Month*1),
	newFood("Сожри карбонару", fr.Month*2),
	newFood("Сожри торт", fr.Month*1),
}

var chooser = newChooser()

func Jri() string {
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
