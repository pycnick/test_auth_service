package usecase

import (
	"github.com/pycnick/test_auth_service/internal/models"
	"github.com/pycnick/test_auth_service/internal/users"
	"log"
)

type UserUseCase struct {
	uR users.Repository
}

func NewUserUseCase(uR users.Repository) *UserUseCase{
	return &UserUseCase{
		uR: uR,
	}
}

func (uUC *UserUseCase) SignUp(user *models.User) error {
	log.Println(user)
	_, err := uUC.uR.Read(user.Login)

	if err == nil {
		return UserAlreadyExistsError
	}

	user.HashPassword();
	err = uUC.uR.Create(user)

	return err
}