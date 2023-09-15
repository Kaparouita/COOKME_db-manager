package main

import (
	"log"

	"db-manager/core"
	"db-manager/repositories"

	"github.com/Kaparouita/models/models"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := repositories.NewDbRepo()
	srv := core.NewService(db)
	srv.GetUser(&models.User{})
	//handler := handlers.NewHandler(srv)
	//server := server.NewService(handler)

	//server.Initialize()

}
