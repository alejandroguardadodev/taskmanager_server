package routes

import (
	"github.com/gofiber/fiber/v2"
	"taskmanagerserver.com/api/controllers"
	"taskmanagerserver.com/api/midlewares"
)

func registerUserRoutes(api fiber.Router) {
	authRoutes := api.Group("/auth").Use(midlewares.RouteRequestToJSON)

	authRoutes.Post("/signup", controllers.AuthRegister)
	authRoutes.Post("/login", controllers.AuthLogin)
}

func Register(app *fiber.App) {
	api := app.Group("/api")

	registerUserRoutes(api)
}
