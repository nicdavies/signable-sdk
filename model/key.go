package model

import (
	"github.com/bxcodec/faker/v3"
)

type Key struct {
	Token string `faker:"uuid_digit"`
	Type  string `faker:"-"`
}

var Keys []Key

// InitKey sets up the data fixtures for api keys
func InitKey() ([]Key, error) {
	successKey := Key{}
	errorKey := Key{}

	err := faker.FakeData(&successKey)
	if err != nil {
		return Keys, err
	}

	err = faker.FakeData(&errorKey)
	if err != nil {
		return Keys, err
	}

	// just need to set the Type on both keys
	successKey.Type = "success"
	errorKey.Type = "error"

	Keys = append(Keys, successKey, errorKey)

	return Keys, nil
}
