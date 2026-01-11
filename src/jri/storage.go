package jri

import (
	"errors"
	"fmt"
)

const (
	BasedPresetId    = "based"
	ThaiPresetId     = "thai"
	GeorgianPresetId = "georgian"
	defaultId        = BasedPresetId
)

var Presets = map[string]Preset{
	BasedPresetId:    BasedPreset,
	ThaiPresetId:     ThaiPreset,
	GeorgianPresetId: GeorgianPreset,
}

type Storage struct {
	// UserId to presetId
	data map[int64]string
}

func NewStorage() *Storage {
	return &Storage{data: make(map[int64]string)}
}

func (s *Storage) PutUserPreset(userid int64, presetId string) error {
	if s.data == nil {
		return errors.New("storage is not initialized")
	}

	s.data[userid] = presetId
	return nil
}

func (s *Storage) GetOrInitUserPreset(userid int64) (string, error) {
	if s.data == nil {
		return "", errors.New("storage is not initialized")
	}

	presetId, ok := s.data[userid]
	if !ok || presetId == "" {
		err := s.PutUserPreset(userid, defaultId)
		if err != nil {
			return "", fmt.Errorf("failed to init default user preset: %w", err)
		} else {
			return defaultId, nil
		}
	}

	return presetId, nil
}
