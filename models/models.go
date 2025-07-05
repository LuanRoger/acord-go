package models

import "time"

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Success bool   `json:"success"`
}

type Activity struct {
	State, Details                                                              string
	StartTimestamp, EndTimestamp                                                *time.Time
	LargeImageKey, SmallImageKey, LargeImageText, SmallImageText, DominantColor *string
}

type ActivityMetadata struct {
	PreviousActivity                               *Activity
	LastActivityUpdateTimestamp                    *int64
	LastActivityUpdateDate, LastAccessActivityDate *time.Time
}

type GetActivityResponse struct {
	State          string    `json:"state"`
	Details        string    `json:"details"`
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
	LargeImageKey  string    `json:"largeImageKey"`
	SmallImageKey  string    `json:"smallImageKey"`
	LargeImageText string    `json:"largeImageText"`
	SmallImageText string    `json:"smallImageText"`
	DominantColor  string    `json:"dominantColor"`
}

type PostNewActivity struct {
	State          string `json:"state"`
	Details        string `json:"details"`
	StartTimestamp int64  `json:"startTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp"`
	LargeImageKey  string `json:"largeImageKey"`
	SmallImageKey  string `json:"smallImageKey"`
	LargeImageText string `json:"largeImageText"`
	SmallImageText string `json:"smallImageText"`
}
