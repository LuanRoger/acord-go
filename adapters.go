package main

import (
	"acord-go/models"
	"time"
)

func PostNewActivityToActivity(value models.PostNewActivity) *models.Activity {
	var startTimestamp, endTimestamp = time.Unix(value.StartTimestamp, 0), time.Unix(value.EndTimestamp, 0)

	return &models.Activity{
		State:          value.State,
		Details:        value.Details,
		StartTimestamp: &startTimestamp,
		EndTimestamp:   &endTimestamp,
		LargeImageKey:  &value.LargeImageKey,
		SmallImageKey:  &value.SmallImageKey,
		LargeImageText: &value.LargeImageText,
		SmallImageText: &value.SmallImageText,
	}
}
