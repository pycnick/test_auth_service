package users

import "github.com/pycnick/test_auth_service/internal/models"

type UseCase interface {
	SignUp(user *models.User) error
}
