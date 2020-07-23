package database

import (
	"github.com/jackc/pgx"
	"github.com/kr/pretty"
	"log"
	"os"
	"time"
)

type Database struct {
	Conn *pgx.ConnPool
}

type meta struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func NewDatabase() (*Database, error) {
	db := meta{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PSWD"),
	}

	connString := pretty.Sprintf("postgresql://%s:%s@%s:%s/%s",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Database)

	config, err := pgx.ParseConnectionString(connString)

	if err != nil {
		log.Fatal(DatabaseMetaError)
	}

	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 100,
	})

	reconnection := 5
	for err != nil && reconnection > 0 {
		conn, err = pgx.NewConnPool(pgx.ConnPoolConfig{
			ConnConfig:     config,
			MaxConnections: 100,
		})
		time.Sleep(time.Millisecond * 200)
		reconnection--
		log.Println(err)
	}

	if reconnection == 0 {
		return nil, DatabaseConnectionError
	}

	return &Database{
		Conn: conn,
	}, nil
}


