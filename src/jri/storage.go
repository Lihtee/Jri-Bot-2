package jri

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

type Storage interface {
	Close() error
	PutUserPreset(userid int64, presetId string) error
	GetOrInitUserPreset(userid int64) (string, error)
}

type storage struct {
	db *bolt.DB
}

var _ Storage = (*storage)(nil)

const (
	BasedPresetId    = "based"
	ThaiPresetId     = "thai"
	GeorgianPresetId = "georgian"
	DefaultPresetId  = BasedPresetId
)

type BucketName string

const (
	UserPresets BucketName = "userPresets"
)

var Presets = map[string]Preset{
	BasedPresetId:    BasedPreset,
	ThaiPresetId:     ThaiPreset,
	GeorgianPresetId: GeorgianPreset,
}

func NewStorage(filepath string) Storage {
	db, err := bolt.Open(filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &storage{db: db}
}

func (s *storage) Close() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("could not close database: %v", err)
	}

	return nil
}

func (s *storage) PutUserPreset(userid int64, presetId string) error {
	db := s.db
	if db == nil {
		return errors.New("storage is not initialized")
	}

	err := db.Update(func(tx *bolt.Tx) error {
		userKey := []byte(strconv.FormatInt(userid, 10))
		return putUserPreset(tx, userKey, presetId)
	})
	if err != nil {
		return fmt.Errorf("could not put user preset: %v", err)
	}

	return nil
}

func putUserPreset(tx *bolt.Tx, userid []byte, presetId string) error {
	bucket, err := tx.CreateBucketIfNotExists([]byte(UserPresets))
	if err != nil {
		return err
	}

	err = bucket.Put(userid, []byte(presetId))
	if err != nil {
		return err
	}
	return nil
}

func getUserPreset(tx *bolt.Tx, userid []byte) (bool, string, error) {
	bucket, err := tx.CreateBucketIfNotExists([]byte(UserPresets))
	if err != nil {
		return false, "", err
	}

	value := bucket.Get(userid)
	if value == nil {
		return false, "", nil
	} else {
		return true, string(value), nil
	}
}

func (s *storage) GetOrInitUserPreset(userid int64) (string, error) {
	db := s.db
	if db == nil {
		return "", errors.New("storage is not initialized")
	}

	var presetId string
	err := db.Update(func(tx *bolt.Tx) error {
		userKey := []byte(strconv.FormatInt(userid, 10))
		ok, value, err := getUserPreset(tx, userKey)
		if ok {
			presetId = value
			return nil
		}

		if err != nil {
			return err
		}

		err = putUserPreset(tx, userKey, DefaultPresetId)
		if err != nil {
			return err
		} else {
			presetId = DefaultPresetId
			return nil
		}
	})

	return presetId, err
}
