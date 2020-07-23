package models

import (
	"errors"
)

var (
	LoginFormatError = errors.New("wrong login format")
	EmailFormatError = errors.New("wrong email format")
	PasswordFormatError = errors.New("wrong password format")
	PhoneFormatError = errors.New("wrong phone number format")
)
