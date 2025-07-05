package utils

import "acord-go/models"

func EmptyResponse(message string, success bool) models.ApiResponse {
	return models.ApiResponse{
		Message: message,
		Success: success,
	}
}
