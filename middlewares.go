package main

import (
	"acord-go/utils"
	"os"

	"github.com/gofiber/fiber/v3"
)

func AuthToken(c fiber.Ctx) error {
	var headers = c.Get("Authorization")
	var authValueEnv = os.Getenv("AUTH_TOKEN")

	if authValueEnv != headers {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(utils.EmptyResponse("Unauthorized", false))
	}

	return c.Next()
}
