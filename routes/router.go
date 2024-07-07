package routes

import (
	"github.com/gofiber/fiber/v2"
	"taskmanagerserver.com/api/controllers"
)

func registerUserRoutes(api fiber.Router) {
	authRoutes := api.Group("/auth")

	authRoutes.Post("/", controllers.AuthRegister)
}

func Register(app *fiber.App) {
	api := app.Group("/api")

	registerUserRoutes(api)
}
