package main

import (
	"github.com/devghor/go-starter/internal/database"
	"github.com/devghor/go-starter/internal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
