package usecase

import "errors"

var (
	UserAlreadyExistsError = errors.New("user with the same login already exists")
	CreateUserError = errors.New("error while creating user")
)
