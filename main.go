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

	var activityGroup = app.Group("/activity")

	activityGroup.Get("/", ActivityGetHandler)
	activityGroup.Post("/", ActivityPostHandler)
	activityGroup.Delete("/", ActivityDeleteHandler)

	app.Listen(":3000")
}
