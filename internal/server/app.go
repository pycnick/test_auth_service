package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/pycnick/test_auth_service/internal/database"
	"github.com/pycnick/test_auth_service/internal/sessions"
	_sessionsDelivery "github.com/pycnick/test_auth_service/internal/sessions/delivery"
	_sessionsRepository "github.com/pycnick/test_auth_service/internal/sessions/repository"
	_sessionsUseCase "github.com/pycnick/test_auth_service/internal/sessions/usecase"
	"github.com/pycnick/test_auth_service/internal/users"
	_usersDelivery "github.com/pycnick/test_auth_service/internal/users/delivery"
	_usersRepository "github.com/pycnick/test_auth_service/internal/users/repository"
	_usersUseCase "github.com/pycnick/test_auth_service/internal/users/usecase"
	"log"
)

type App struct {
	router     *gin.Engine
	db         *pgx.ConnPool
	usersUC    users.UseCase
	sessionsUC sessions.UseCase
}

func NewApp() *App {
	db, err := database.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	uR := _usersRepository.NewUserRepository(db.Conn)
	uUC := _usersUseCase.NewUserUseCase(uR)

	sR := _sessionsRepository.NewSessionRepository(db.Conn)
	sUC := _sessionsUseCase.NewSessionUseCase(sR, uR)

	return &App{
		router:     gin.Default(),
		db: db.Conn,
		usersUC:    uUC,
		sessionsUC: sUC,
	}
}

func (a *App) Run() {
	defer a.db.Close()

	api := a.router.Group("/api/v1")
	_usersDelivery.NewUserHandlers(api, a.usersUC)
	_sessionsDelivery.NewSessionHandlers(api, a.sessionsUC)

	log.Fatal(a.router.Run(":8080"))
}
