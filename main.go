package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	//platform.ConnectDB()
	//
	//routes.SetupRoutes(app)
	//log.Fatal(app.Listen(":3000"))
}
