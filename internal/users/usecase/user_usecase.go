package usecase

import (
	"github.com/pycnick/test_auth_service/internal/models"
	"github.com/pycnick/test_auth_service/internal/users"
)

type UserUseCase struct {
	userR users.Repository
}

func NewUserUseCase(uR users.Repository) *UserUseCase{
	return &UserUseCase{
		userR: uR,
	}
}

func (uUC *UserUseCase) SignUp(user *models.User) error {
	return nil
}