package server

import (
	"db-manager/ports"

	"github.com/Kaparouita/models/rabbitmq"
	"github.com/streadway/amqp"
	// Other necessary imports
)

type Server struct {
	IngredientHandler ports.IngredentHandler
	RecipeHandler     ports.RecipeHandler
	UserHandler       ports.UserHandler
}

func NewService(ingredientHandler ports.IngredentHandler, recipeHandler ports.RecipeHandler, userHandler ports.UserHandler) *Server {
	return &Server{
		ingredientHandler,
		recipeHandler,
		userHandler,
	}
}

func (server *Server) DeclareQueues(ch *amqp.Channel) error {
	// Declare queues here
	// Example:
	_, err := ch.QueueDeclare("SaveRecipeQueue", false, false, false, false, nil)
	rabbitmq.RegisterConsumer(server.RecipeHandler.SaveRecipe, ch, "SaveRecipeQueue", "SaveRecipeQueue")
	if err != nil {
		return err
	}
	// Declare other queues
	return nil
}
