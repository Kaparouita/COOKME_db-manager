package ports

import (
	"github.com/Kaparouita/db-manager/domain"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	InitFunction() error
	GetIngridients(pagin *domain.IngridientPagin) (*domain.IngridientResponse, error)
	SaveIngridient(Ingridient *domain.Ingridient) *domain.Response
	UpdateIngridient(Ingridient *domain.Ingridient) *domain.Response
	GetIngridient(Ingridient *domain.Ingridient) (*domain.Ingridient, *domain.Response)
	DeleteIngridient(Ingridient *domain.Ingridient) *domain.Response
	GetIngridientsPerUni(group string) ([]domain.IngridientsPerUni, *domain.Response)

	Register(user *domain.User) *domain.User
	GetUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) *domain.User
	GetUser(user *domain.User) *domain.User
	DeleteUser(user *domain.User) *domain.Response
	ReviewIngridient(review *domain.Review) *domain.Response
	GetUserPerType(group string) ([]domain.UserPerType, *domain.Response)

	Login(login *domain.LoginResp) *domain.User
}

type Db interface {
	SaveIngridient(Ingridient *domain.Ingridient) error
	GetIngridients(pagin *domain.IngridientPagin) (*domain.IngridientResponse, error)
	UpdateIngridient(Ingridient *domain.Ingridient) error
	GetIngridient(Ingridient *domain.Ingridient) error
	DeleteIngridient(Ingridient *domain.Ingridient) error
	GetIngridientsPerUni(groupField string) ([]domain.IngridientsPerUni, error)

	SaveUser(user *domain.User) error
	GetUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) error
	GetUser(user *domain.User) error
	DeleteUser(user *domain.User) error
	GetUserPerUserType(groupField string) ([]domain.UserPerType, error)

	Login(login *domain.LoginResp) (*domain.User, error)
}

type Handler interface {
	SaveIngridient(c *fiber.Ctx) error
	UpdateIngridient(c *fiber.Ctx) error
	GetIngridient(c *fiber.Ctx) error
	GetIngridients(c *fiber.Ctx) error
	DeleteIngridient(c *fiber.Ctx) error
	GetIngridientsPerUni(c *fiber.Ctx) error

	Register(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	ReviewIngridient(c *fiber.Ctx) error
	GetUsersPerType(c *fiber.Ctx) error

	Login(c *fiber.Ctx) error
}
