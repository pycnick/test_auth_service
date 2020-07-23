package usecase

import "errors"

var (
	UserDoesNotExistsError = errors.New("user with the same login does not exist")
	WrongPasswordError = errors.New("wrong password")
)
