package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
	"github.com/pycnick/test_auth_service/internal/database"
	"github.com/pycnick/test_auth_service/internal/sessions"
	"github.com/pycnick/test_auth_service/internal/users"
	_usersUseCase "github.com/pycnick/test_auth_service/internal/users/usecase"
	_usersRepository "github.com/pycnick/test_auth_service/internal/users/repository"
	"log"
)

type App struct {
	router *gin.Engine
	usersUC users.UseCase
	sessionsUC sessions.UseCase
}

func NewApp() *App {
	db, err := database.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	uR := _usersRepository.NewUserRepository(db.Conn)
	uUC := _usersUseCase.NewUserUseCase(uR)

	return &App{
		router:     gin.Default(),
		usersUC:    uUC,
		sessionsUC: nil,
	}
}

func (a *App) Run() {

}