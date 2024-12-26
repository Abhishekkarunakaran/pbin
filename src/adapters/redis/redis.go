package redis

import "errors"

var (
	ErrSaveData         = errors.New("failed to save data in db")
	ErrSerializeData    = errors.New("failed to serialize data")
	ErrExpire           = errors.New("failed to set expire for the data")
	ErrValueDoesntExist = errors.New("value doesn't exists in the db")
	ErrFetchValue       = errors.New("failed to fetch data from db")
	ErrDeserializeData  = errors.New("failed to deserialize data")
	ErrRemove           = errors.New("failed to delete the data")
)
