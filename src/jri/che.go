package jri

import (
	"fmt"
	"log"
	"unicode"

	fr "github.com/dreyspi/jribot2/frequency"
	wr "github.com/mroth/weightedrand"
)

type Che struct {
	storage Storage
}

func NewChe(storage Storage) *Che {
	if storage == nil {
		log.Fatal("storage is nil")
	}
	return &Che{storage: storage}
}

func (che *Che) Sojrat(userId int64) (string, error) {
	presetId, err := che.storage.GetOrInitUserPreset(userId)
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

func (che *Che) UMenya(userId int64) (string, error) {
	presetId, err := che.storage.GetOrInitUserPreset(userId)
	if err != nil {
		return "", fmt.Errorf("failed to get food for user %d: %w", userId, err)
	}

	return presetId, nil
}

func (che *Che) Viberi(userId int64, presetId string) error {
	return che.storage.PutUserPreset(userId, presetId)
}

func Tut(packId string) ([]string, error) {
	if packId == "" {
		return nil, fmt.Errorf("empty pack id")
	}

	preset, ok := Presets[packId]
	if !ok {
		return nil, fmt.Errorf("no preset found for pack id %s", packId)
	}

	result := make([]string, len(preset))
	for index, food := range preset {
		result[index] = formatCheTut(food)
	}

	return result, nil
}

func formatCheTut(food *Food) string {
	foodName := []rune(food.NominativeName)
	foodName[0] = unicode.ToUpper(foodName[0])

	period := formatCheTutPeriod(food)

	return fmt.Sprintf("%s *%s* \\-\\> %s", food.Emoji, string(foodName), period)
}

func formatCheTutPeriod(food *Food) string {
	if food.Factor == 0 {
		return "*не жрать*"
	}

	var frequency string
	switch food.Frequency {
	case fr.Year:
		frequency = "год"
		break
	case fr.Month:
		frequency = "месяц"
		break
	case fr.Week:
		frequency = "неделю"
		break
	case fr.Day:
		frequency = "день"
		break
	default:
		return "*не жрать*"
	}

	var times string
	if food.Factor == 1 {
		times = "раз"
	} else if food.Factor >= 2 && food.Factor <= 4 {
		times = fmt.Sprintf("%d раза", food.Factor)
	} else {
		times = fmt.Sprintf("%d раз", food.Factor)
	}

	return fmt.Sprintf("*%s* в *%s*", times, frequency)
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
