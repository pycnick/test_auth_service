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
	return nil
}
