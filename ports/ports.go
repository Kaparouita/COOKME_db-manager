package ports

import (
	"github.com/Kaparouita/models/models"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	InitFunction() error
	SaveIngredient(Ingredient *models.Ingredient) *models.Response
	UpdateIngredient(Ingredient *models.Ingredient) *models.Response
	GetIngredient(Ingredient *models.Ingredient) (*models.Ingredient, *models.Response)
	DeleteIngredient(Ingredient *models.Ingredient) *models.Response

	Register(user *models.User) *models.User
	GetUsers() ([]models.User, error)
	UpdateUser(user *models.User) *models.User
	GetUser(user *models.User) *models.User
	DeleteUser(user *models.User) *models.Response

	Login(login *models.LoginResp) *models.User
}

type Db interface {
	SaveIngredient(Ingredient *models.Ingredient) error
	UpdateIngredient(Ingredient *models.Ingredient) error
	GetIngredient(Ingredient *models.Ingredient) error
	GetIngredients() ([]models.Ingredient, error)
	DeleteIngredient(Ingredient *models.Ingredient) error

	SaveMarketIngredient(MarketIngredient *models.MarketIngredient) error
	UpdateMarketIngredient(MarketIngredient *models.MarketIngredient) error
	GetMarketIngredient(MarketIngredient *models.MarketIngredient) error
	GetMarketIngredients() ([]models.MarketIngredient, error)
	DeleteMarketIngredient(MarketIngredient *models.MarketIngredient) error

	SaveRecipe(Recipe *models.Recipe) error
	UpdateRecipe(Recipe *models.Recipe) error
	GetRecipe(Recipe *models.Recipe) error
	GetRecipes() ([]models.Recipe, error)
	DeleteRecipe(Recipe *models.Recipe) error

	SaveFinalRecipe(FinalRecipe *models.FinalRecipe) error
	UpdateFinalRecipe(FinalRecipe *models.FinalRecipe) error
	GetFinalRecipe(FinalRecipe *models.FinalRecipe) error
	GetFinalRecipes() ([]models.FinalRecipe, error)
	DeleteFinalRecipe(FinalRecipe *models.FinalRecipe) error

	SaveUser(user *models.User) error
	GetUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	GetUser(user *models.User) error
	DeleteUser(user *models.User) error

	Login(login *models.LoginResp) (*models.User, error)
}

type Handler interface {
	SaveIngredient(c *fiber.Ctx) error
	UpdateIngredient(c *fiber.Ctx) error
	GetIngredient(c *fiber.Ctx) error
	GetIngredients(c *fiber.Ctx) error
	DeleteIngredient(c *fiber.Ctx) error

	Register(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error

	Login(c *fiber.Ctx) error
}
