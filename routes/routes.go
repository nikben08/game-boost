package routes

import (
	"game-boost/handlers"

	"github.com/gofiber/fiber/v2"
)

func Initalize(app *fiber.App) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/signup", handlers.Signup)
	auth.Post("/signin", handlers.Signin)

	api.Post("/order", handlers.CreateOrder)
	api.Put("/order/:id", handlers.EditOrder)
	api.Delete("/order/:id", handlers.DeleteOrder)
	api.Get("/order", handlers.GetAllOrders)
	api.Get("/order/:id", handlers.GetOrder)
}
