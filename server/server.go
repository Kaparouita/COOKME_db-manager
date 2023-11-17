package server

import (
	"db-manager/ports"
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
