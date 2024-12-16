package service

import "errors"

var (
	ErrGenUUID  = errors.New("failed to generate unique id")
	ErrHashPass = errors.New("failed to hash the password")
	ErrGenBlock = errors.New("failed to create cipher block")
	ErrGenGCM   = errors.New("failed to generate gcm for the ciphe block")
	ErrGenNonce = errors.New("failed to generate random nonce")
	ErrSaveData = errors.New("failed to save the data")
	ErrGetData = errors.New("failed to get data from db")
	ErrIncorrectPassword = errors.New("password hash and entered password is not matching")
	ErrDecrypting = errors.New("failed to decrypt the content")
)
