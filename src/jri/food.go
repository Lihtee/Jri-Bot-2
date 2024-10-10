package jri

import (
	"fmt"
	wr "github.com/mroth/weightedrand"
)

type Food struct {
	Name   string
	Weight int
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

func (food *Food) toChoice() wr.Choice {
	return wr.Choice{Item: food.Name, Weight: uint(food.Weight)}
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

func NewFood(name string, weight int) *Food {
	return &Food{name, weight}
}
