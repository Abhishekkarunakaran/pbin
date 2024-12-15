package redis

import "errors"

var (
	ErrSaveData = errors.New("failed to save data in db")
	ErrSerializeData = errors.New("failed to serialize data")
	ErrExpire = errors.New("failed to set expire for the data")
)
