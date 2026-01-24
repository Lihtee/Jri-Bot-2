package jri

import (
	fr "github.com/dreyspi/jribot2/frequency"
)

type Food struct {
	NominativeName string
	GenitiveName   string
	Emoji          string
	Frequency      fr.Frequency
	Factor         int
}

func NewFood(nominativeName string, genitiveName string, emoji string, frequency fr.Frequency, factor int) *Food {
	return &Food{nominativeName, genitiveName, emoji, frequency, factor}
}
