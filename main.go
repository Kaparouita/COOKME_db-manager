package main

import (
	"db-manager/core"
	"db-manager/handlers"
	"db-manager/repositories"
	"db-manager/server"
	"fmt"
	"log"
	"os"

	"github.com/Kaparouita/models/models"
	"github.com/Kaparouita/models/myrabbit/amqp"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handler := amqp.AmqpHandlerInstance(os.Getenv("RABBITMQ_DIAL"))

	rc := make(chan bool)
	handler.AutoRedial(os.Getenv("RABBITMQ_DIAL"), rc)

	fmt.Println("Start server")

	// go func() {
	// 	for ; true; <-rc {
	db := repositories.NewDbRepo()
	recipeService := core.NewRecipeService(db)
	ingredientService := core.NewIngredientService(db)
	userService := core.NewUserService(db)

	recipeHandler := handlers.NewRecipeHandler(recipeService, handler)
	ingedientHandler := handlers.NewIngredientHandler(ingredientService, handler)
	userHandler := handlers.NewUserHandler(userService, handler)

	server := server.NewServer(ingedientHandler, recipeHandler, userHandler)
	server.Initialize(handler)

	user := &models.User{
		Username: "test",
		Password: "test",
	}
	userService.Register(user)
	// 	}
	// }()

	forever := make(chan bool)

	<-forever
}
