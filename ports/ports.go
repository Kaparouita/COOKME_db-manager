package ports

import (
	"github.com/Kaparouita/models/models"
	"github.com/Kaparouita/models/myrabbit"
)

type IngredientService interface {
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
	SaveMarketIngredient(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	UpdateMarketIngredient(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	GetMarketIngredient(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	GetMarketIngredients(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	DeleteMarketIngredient(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
}

type RecipeHandler interface {
	SaveRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	UpdateRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	GetRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	GetRecipes(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	DeleteRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)

	SaveFinalRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	UpdateFinalRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	GetFinalRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	GetFinalRecipes(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	DeleteFinalRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
}

type UserHandler interface {
	Register(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	UpdateUser(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	GetUser(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	GetUsers(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	DeleteUser(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
	Login(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel)
}
