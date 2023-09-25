package ports

import (
	"github.com/Kaparouita/models/models"
	"github.com/streadway/amqp"
)

type IngredientService interface {
	SaveIngredient(Ingredient *models.Ingredient) *models.Response
	UpdateIngredient(Ingredient *models.Ingredient) *models.Response
	GetIngredient(Ingredient *models.Ingredient) (*models.Ingredient, *models.Response)
	DeleteIngredient(Ingredient *models.Ingredient) *models.Response
	GetIngredients() ([]models.Ingredient, error)

	SaveMarketIngredient(MarketIngredient *models.MarketIngredient) *models.Response
	UpdateMarketIngredient(MarketIngredient *models.MarketIngredient) *models.Response
	GetMarketIngredient(MarketIngredient *models.MarketIngredient) (*models.MarketIngredient, *models.Response)
	DeleteMarketIngredient(MarketIngredient *models.MarketIngredient) *models.Response
	GetMarketIngredients() ([]models.MarketIngredient, error)
}

type RecipeService interface {
	SaveRecipe(Recipe *models.Recipe) *models.Response
	UpdateRecipe(Recipe *models.Recipe) *models.Response
	GetRecipe(Recipe *models.Recipe) (*models.Recipe, *models.Response)
	DeleteRecipe(Recipe *models.Recipe) *models.Response
	GetRecipes() ([]models.Recipe, error)

	SaveFinalRecipe(FinalRecipe *models.FinalRecipe) *models.Response
	UpdateFinalRecipe(FinalRecipe *models.FinalRecipe) *models.Response
	GetFinalRecipe(FinalRecipe *models.FinalRecipe) (*models.FinalRecipe, *models.Response)
	DeleteFinalRecipe(FinalRecipe *models.FinalRecipe) *models.Response
	GetFinalRecipes() ([]models.FinalRecipe, error)
}

type UserService interface {
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

type IngredentHandler interface {
	SaveIngredient(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	UpdateIngredient(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetIngredient(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetIngredients(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	DeleteIngredient(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error

	SaveMarketIngredient(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	UpdateMarketIngredient(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetMarketIngredient(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetMarketIngredients(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	DeleteMarketIngredient(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
}

type RecipeHandler interface {
	SaveRecipe(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	UpdateRecipe(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetRecipe(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetRecipes(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	DeleteRecipe(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error

	SaveFinalRecipe(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	UpdateFinalRecipe(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetFinalRecipe(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetFinalRecipes(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	DeleteFinalRecipe(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
}

type UserHandler interface {
	Register(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	UpdateUser(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetUser(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	GetUsers(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
	DeleteUser(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error

	Login(msgs <-chan amqp.Delivery, pubCh amqp.Channel, subCh amqp.Channel) error
}
