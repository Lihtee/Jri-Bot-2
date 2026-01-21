package jri

import (
	"fmt"
	"unicode"

	fr "github.com/dreyspi/jribot2/frequency"
	wr "github.com/mroth/weightedrand"
)

type Food struct {
	NominativeName string
	GenitiveName   string
	Emoji          string
	Frequency      fr.Frequency
	Factor         int
}

var storage = NewStorage()

func Jri(userId int64) (string, error) {
	presetId, err := storage.GetOrInitUserPreset(userId)
	if err != nil {
		return "", fmt.Errorf("failed to get food for user %d: %w", userId, err)
	}

	preset := Presets[presetId]
	chooser, err := newChooser(preset)
	if err != nil {
		return "", err
	}

	return chooser.Pick().(string), nil
}

func Eda(userId int64) (string, error) {
	presetId, err := storage.GetOrInitUserPreset(userId)
	if err != nil {
		return "", fmt.Errorf("failed to get food for user %d: %w", userId, err)
	}

	return presetId, nil
}

func SetEda(userId int64, eda string) error {
	return storage.PutUserPreset(userId, eda)
}

func CheTut(packId string) ([]string, error) {
	if packId == "" {
		return nil, fmt.Errorf("empty pack id")
	}

	preset, ok := Presets[packId]
	if !ok {
		return nil, fmt.Errorf("no preset found for pack id %s", packId)
	}

	result := make([]string, len(preset))
	for index, food := range preset {
		runes := []rune(food.NominativeName)
		runes[0] = unicode.ToUpper(runes[0])
		result[index] = fmt.Sprintf("%s %s", food.Emoji, string(runes))
	}

	return result, nil
}

func (food *Food) toChoice() wr.Choice {
	return wr.Choice{
		Item:   buildChoiceName(food),
		Weight: uint(int(food.Frequency) * food.Factor),
	}
}

func buildChoiceName(food *Food) string {
	if len(food.GenitiveName) == 0 {
		return "Ниче не жри"
	}

	return fmt.Sprintf("Сожри %s %s", food.GenitiveName, food.Emoji)
}

func newChooser(preset Preset) (*wr.Chooser, error) {
	choices := []wr.Choice{}
	for _, food := range preset {
		choices = append(choices, food.toChoice())
	}
	chooser, err := wr.NewChooser(choices...)
	if err != nil {
		return nil, fmt.Errorf("failed to init choser: %s", err)
	}

	return chooser, nil
}

func NewFood(nominativeName string, genitiveName string, emoji string, frequency fr.Frequency, factor int) *Food {
	return &Food{nominativeName, genitiveName, emoji, frequency, factor}
}
