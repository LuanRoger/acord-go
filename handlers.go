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

	var actvityResponse = ActivityToGetActivityResponse(*ActivityInstance)
	var response = models.ApiResponse{
		Message: "Activity found",
		Data:    actvityResponse,
		Success: true,
	}
	c.Status(fiber.StatusOK)
	return c.JSON(response)
}

func ActivityPostHandler(c fiber.Ctx) error {
	var bodyData = new(models.PostNewActivity)
	bindError := c.Bind().Body(bodyData)
	if bindError != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(utils.EmptyResponse("Invalid Body", false))
	}

	var newActivity, adapterError = PostNewActivityToActivity(*bodyData)
	if adapterError != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(utils.EmptyResponse(adapterError.Error(), false))
	}

	ActivityInstance = newActivity

	var actvityResponse = ActivityToGetActivityResponse(*ActivityInstance)
	var response = models.ApiResponse{
		Message: "Activity created",
		Data:    actvityResponse,
		Success: true,
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(response)
}

func ActivityDeleteHandler(c fiber.Ctx) error {
	ActivityInstance = nil

	var response = models.ApiResponse{
		Message: "Activity deleted",
		Success: true,
	}
	c.Status(fiber.StatusOK)
	return c.JSON(response)
}
