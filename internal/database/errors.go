package database

import "errors"

var (
	DatabaseMetaError = errors.New("Wrong database config data")
	DatabaseConnectionError = errors.New("no such database connection")
)
