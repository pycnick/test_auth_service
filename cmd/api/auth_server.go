package main

import "github.com/pycnick/test_auth_service/internal/server"

func main() {
	app := server.NewApp()
	app.Run()
}
