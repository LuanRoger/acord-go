package main

import (
	"acord-go/models"
	"acord-go/utils"
	"time"

	"github.com/cenkalti/dominantcolor"
)

func PostNewActivityToActivity(value models.PostNewActivity) (*models.Activity, error) {
	var startTimestamp, endTimestamp = time.Unix(value.StartTimestamp, 0), time.Unix(value.EndTimestamp, 0)
	var image, err = utils.DownloadRemoteImage(value.SmallImageKey)
	if err != nil {
		return nil, err
	}
	var dominantColor = dominantcolor.Hex(dominantcolor.Find(*image))

	return &models.Activity{
		State:          value.State,
		Details:        value.Details,
		StartTimestamp: &startTimestamp,
		EndTimestamp:   &endTimestamp,
		LargeImageKey:  &value.LargeImageKey,
		SmallImageKey:  &value.SmallImageKey,
		LargeImageText: &value.LargeImageText,
		SmallImageText: &value.SmallImageText,
		DominantColor:  &dominantColor,
	}, nil
}

func ActivityToGetActivityResponse(value models.Activity) *models.GetActivityResponse {
	return &models.GetActivityResponse{
		State:          value.State,
		Details:        value.Details,
		StartTimestamp: *value.StartTimestamp,
		EndTimestamp:   *value.EndTimestamp,
		LargeImageKey:  *value.LargeImageKey,
		SmallImageKey:  *value.SmallImageKey,
		LargeImageText: *value.LargeImageText,
		SmallImageText: *value.SmallImageText,
		DominantColor:  *value.DominantColor,
	}
}
