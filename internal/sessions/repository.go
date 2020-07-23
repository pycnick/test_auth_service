package sessions

import "github.com/pycnick/test_auth_service/internal/models"

type Repository interface {
	ReadByUser(login string) (*models.Session, error)
	ReadByToken(token string) (*models.Session, error)
	Create(session *models.Session) error
}
