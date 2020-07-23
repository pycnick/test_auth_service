package sessions

import "github.com/pycnick/test_auth_service/internal/models"

type UseCase interface {
	SignIn(login string, password string) (*models.Session, error)
	GetSession(token string) (*models.Session, error)
}
