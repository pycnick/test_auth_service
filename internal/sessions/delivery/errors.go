package delivery

import "errors"

var (
	AlreadyAuthenticatedError = errors.New("already authenticated")
	BadRequestError = errors.New("bad request")
	SignUpError = errors.New("error while signing up")
)
