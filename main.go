package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var app = fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(AuthToken)

	app.Get("/activity", ActivityGetHandler)
	app.Post("/activity", ActivityPostHandler)

	app.Listen(":3000")
}
