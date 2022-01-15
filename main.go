package main

import (
	"api-fiber/pkg/configs"
	"api-fiber/pkg/middleware"
	"api-fiber/pkg/routes"
	"api-fiber/pkg/utils"
	"api-fiber/platform/database"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	n := os.Args[1]
	switch n {
	case "migrate":
		database.Migrate()
	default:
		// Define Fiber config.
		config := configs.FiberConfig()

		// Define a new Fiber app with config.
		app := fiber.New(config)

		// Middlewares.
		middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

		// Routes.
		//routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
		routes.PublicRoutes(app) // Register a public routes for app.
		//routes.PrivateRoutes(app) // Register a private routes for app.
		//routes.NotFoundRoute(app) // Register route for 404 Error.

		// Start server (with or without graceful shutdown).
		if utils.GetEnv("STAGE_STATUS", "dev") == "dev" {
			utils.StartServer(app)
		} else {
			utils.StartServerWithGracefulShutdown(app)
		}
	}
}
