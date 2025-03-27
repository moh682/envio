package main

import (
	"log"

	"github.com/moh682/envio/backend/internal/api"
	"github.com/moh682/envio/backend/internal/domain/config"
	"github.com/moh682/envio/backend/internal/frameworks/postgres"
)

func main() {

	// Connect to the database
	internalConfig, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	username := internalConfig.GetPostgresConfig().Username
	password := internalConfig.GetPostgresConfig().Password
	host := internalConfig.GetPostgresConfig().Host
	dbname := internalConfig.GetPostgresConfig().DBName

	db, err := postgres.Connect(username, password, dbname, host)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new HTTP server
	api := api.NewHttpServer(db)

	log.Println("Server is running on port 8080")

	if err := api.ListenAndServe(8080); err != nil {
		panic(err)
	}
}
