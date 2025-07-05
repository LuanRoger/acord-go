package main

import (
	"acord-go/models"
	"acord-go/utils"

	"github.com/gofiber/fiber/v3"
)

func ActivityGetHandler(c fiber.Ctx) error {
	if ActivityInstance == nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(utils.EmptyResponse("Activity is empty", false))
	}

	var response = models.ApiResponse{
		Message: "Activity found",
		Data:    ActivityInstance,
		Success: true,
	}
	c.Status(fiber.StatusOK)
	return c.JSON(response)
}

func ActivityPostHandler(c fiber.Ctx) error {
	var bodyData = new(models.PostNewActivity)
	err := c.Bind().Body(bodyData)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(utils.EmptyResponse("Invalid Body", false))
	}

	var newActivity = PostNewActivityToActivity(*bodyData)
	ActivityInstance = newActivity

	var response = models.ApiResponse{
		Message: "Activity found",
		Data:    ActivityInstance,
		Success: true,
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(response)
}
