package routes

import (
	"api-fiber/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	userRoute := route.Group("/user")
	userRoute.Post("/sign/up", controllers.UserSignUp) // register a new user
	userRoute.Post("/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens
}
