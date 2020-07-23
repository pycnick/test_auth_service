package repository

import (
	"github.com/jackc/pgx"
	"github.com/pycnick/test_auth_service/internal/models"
)

type UserRepository struct {
	db *pgx.ConnPool
}

func NewUserRepository(db *pgx.ConnPool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (uR *UserRepository) Create(user *models.User) error {
	_, err := uR.db.Exec("insert into users values ($1, $2, $3, $4)",
		user.Login,
		user.Email,
		user.Phone,
		user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (uR *UserRepository) Read(login string) (*models.User, error) {
	user := &models.User{}

	if err := uR.db.QueryRow("select * from users where login = $1",
		login).Scan(&user.Login, &user.Email, &user.Phone, &user.Password); err != nil {
		return nil, err
	}

	return user, nil
}
