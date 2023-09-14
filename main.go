package main

import (
	"github.com/Kaparouita/db-manager/core"
	"github.com/Kaparouita/db-manager/handlers"
	"github.com/Kaparouita/db-manager/repositories"
	"github.com/Kaparouita/db-manager/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := repositories.NewDbRepo()
	srv := core.NewService(db)
	handler := handlers.NewHandler(srv)
	server := server.NewService(handler)

	//initDB with 5 Ingridients
	srv.InitFunction()

	server.Initialize()

}
