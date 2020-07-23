package repository

import (
	"github.com/jackc/pgx"
	"github.com/pycnick/test_auth_service/internal/models"
)

type SessionRepository struct {
	db *pgx.ConnPool
}

func NewSessionRepository(db *pgx.ConnPool) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (sR *SessionRepository) Create(session *models.Session) error {
	_, err := sR.db.Exec("insert into sessions values ($1, $2, $3)",
		session.Token,
		session.UserID,
		session.Expiration)

	if err != nil {
		return err
	}

	return nil
}

func (sR *SessionRepository) ReadByUser(login string) (*models.Session, error) {
	session := &models.Session{}
	if err := sR.db.QueryRow("select * from sessions where user_id = $1",
		login).Scan(
		&session.Token,
		&session.UserID,
		&session.Expiration); err != nil {
		return nil, err
	}

	return session, nil
}

func (sR *SessionRepository) ReadByToken(token string) (*models.Session, error) {
	session := &models.Session{}
	if err := sR.db.QueryRow("select * from sessions where token = $1",
		token).Scan(
		&session.Token,
		&session.UserID,
		&session.Expiration); err != nil {
		return nil, err
	}

	return session, nil
}


