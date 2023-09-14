package server

import (
	"github.com/Kaparouita/db-manager/ports"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	handler ports.Handler
}

func NewService(handler ports.Handler) *Server {
	return &Server{handler: handler}
}

func (server *Server) Initialize() {
	app := fiber.New()
	app.Use(cors.New())

	//TODO:
	//admin := app.Group("/admin")
	//libraryAdmin := app.Group("/libraryAdmin")

	login := app.Group("/login")
	login.Post("/", server.handler.Login)

	users := app.Group("/users")
	users.Get("/", server.handler.GetUsers)
	users.Get("/perType", server.handler.GetUsersPerType)
	users.Post("/", server.handler.Register)
	users.Put("/:id", server.handler.UpdateUser)
	users.Get("/:id", server.handler.GetUser)
	users.Delete("/:id", server.handler.DeleteUser)
	users.Post("/:id/review", server.handler.ReviewIngridient)

	Ingridients := app.Group("/Ingridients")
	Ingridients.Get("/", server.handler.GetIngridients)
	Ingridients.Get("/perUni", server.handler.GetIngridientsPerUni)
	Ingridients.Get("/:id", server.handler.GetIngridient)
	Ingridients.Post("/", server.handler.SaveIngridient)
	Ingridients.Put("/:id", server.handler.UpdateIngridient)
	Ingridients.Delete("/:id", server.handler.DeleteIngridient)

	log.Fatal(app.Listen(":3000"))
}
