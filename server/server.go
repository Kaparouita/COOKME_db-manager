package server

import (
	"db-manager/ports"

	"github.com/Kaparouita/models/myrabbit/amqp"
	// Other necessary imports
)

type Server struct {
	IngredientHandler ports.IngredentHandler
	RecipeHandler     ports.RecipeHandler
	UserHandler       ports.UserHandler
}

func NewServer(ingredientHandler ports.IngredentHandler, recipeHandler ports.RecipeHandler, userHandler ports.UserHandler) *Server {
	return &Server{
		ingredientHandler,
		recipeHandler,
		userHandler,
	}
}

func (server *Server) Initialize(handler *amqp.AmqpHandler) {

	handler.RegisterConsumer("db_manager.get.market_ingedient", server.IngredientHandler.GetMarketIngredient)
	handler.RegisterConsumer("db_manager.get.market_ingedients", server.IngredientHandler.GetMarketIngredients)
	handler.RegisterConsumer("db_manager.save.market_ingedient", server.IngredientHandler.SaveMarketIngredient)
	handler.RegisterConsumer("db_manager.update.market_ingedient", server.IngredientHandler.UpdateMarketIngredient)
	handler.RegisterConsumer("db_manager.delete.market_ingedient", server.IngredientHandler.DeleteMarketIngredient)

	handler.RegisterConsumer("db_manager.get.recipe", server.RecipeHandler.GetRecipe)
	handler.RegisterConsumer("db_manager.get.recipes", server.RecipeHandler.GetRecipes)
	handler.RegisterConsumer("db_manager.save.recipe", server.RecipeHandler.SaveRecipe)
	handler.RegisterConsumer("db_manager.update.recipe", server.RecipeHandler.UpdateRecipe)
	handler.RegisterConsumer("db_manager.delete.recipe", server.RecipeHandler.DeleteRecipe)

	handler.RegisterConsumer("db_manager.get.final_recipe", server.RecipeHandler.GetFinalRecipe)
	handler.RegisterConsumer("db_manager.get.final_recipes", server.RecipeHandler.GetFinalRecipes)
	handler.RegisterConsumer("db_manager.save.final_recipe", server.RecipeHandler.SaveFinalRecipe)
	handler.RegisterConsumer("db_manager.update.final_recipe", server.RecipeHandler.UpdateFinalRecipe)
	handler.RegisterConsumer("db_manager.delete.final_recipe", server.RecipeHandler.DeleteFinalRecipe)

	handler.RegisterConsumer("db_manager.get.user", server.UserHandler.GetUser)
	handler.RegisterConsumer("db_manager.get.users", server.UserHandler.GetUsers)
	handler.RegisterConsumer("db_manager.resgister.user", server.UserHandler.Register)
	handler.RegisterConsumer("db_manager.update.user", server.UserHandler.UpdateUser)
	handler.RegisterConsumer("db_manager.delete.user", server.UserHandler.DeleteUser)

	handler.RegisterConsumer("db_manager.login", server.UserHandler.Login)

}
