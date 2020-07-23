package users

import "github.com/pycnick/test_auth_service/internal/models"

type Repository interface {
	Create(user *models.User) error
	Read(login string) (*models.User, error)
}
