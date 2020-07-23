package usecase

import (
	"github.com/pycnick/test_auth_service/internal/models"
	"github.com/pycnick/test_auth_service/internal/sessions"
	"github.com/pycnick/test_auth_service/internal/users"
)

type SessionUseCase struct {
	sR sessions.Repository
	uR users.Repository
}

func NewSessionUseCase(sR sessions.Repository, uR users.Repository) *SessionUseCase {
	return &SessionUseCase{
		sR: sR,
		uR: uR,
	}
}

func (sUC *SessionUseCase) SignIn(login string, password string) (*models.Session, error) {
	user, err := sUC.uR.Read(login)

	if err != nil {
		return nil, UserDoesNotExistsError
	}

	if !user.VerifyPassword(password) {
		return nil, WrongPasswordError
	}

	session := models.NewSession(login)

	err = sUC.sR.Create(session)

	if err != nil {
		return nil, err
	}

	return session, err
}

func (sUC *SessionUseCase) GetSession(token string) (*models.Session, error) {
	session, err := sUC.sR.ReadByToken(token)
	if err != nil {
		return nil, err
	}

	return session, nil
}
