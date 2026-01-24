package jri

type TestStorage struct {
	data map[int64]string
}

var _ Storage = (*TestStorage)(nil)

func NewTestStorage() Storage {
	return &TestStorage{data: make(map[int64]string)}
}

func (s *TestStorage) Close() error {
	return nil
}

func (s *TestStorage) PutUserPreset(userid int64, presetId string) error {
	s.data[userid] = presetId
	return nil
}

func (s *TestStorage) GetOrInitUserPreset(userid int64) (string, error) {
	if presetId, ok := s.data[userid]; ok {
		return presetId, nil
	}
	s.data[userid] = DefaultPresetId
	return DefaultPresetId, nil
}
