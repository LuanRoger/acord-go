package main

import (
	"acord-go/models"
	"time"
)

var ActivityInstance *models.Activity
var ActivityCleanupTimer *time.Timer

func RestartCleanupTimer() {
	if ActivityCleanupTimer != nil {
		ActivityCleanupTimer.Stop()
	}

	ActivityCleanupTimer = time.AfterFunc(time.Hour, func() {
		ActivityInstance = nil
	})
}
