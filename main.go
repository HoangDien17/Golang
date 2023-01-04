package main

import (
	"employee/app"
	"log"

	"github.com/joho/godotenv"
)

// load .env file
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

// @title My awesome API
// @version 1.0
// @description My awesome API
// @host localhost:4000
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	var server app.Routes
	server.Bootstrap()
}
