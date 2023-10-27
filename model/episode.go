package model

import "time"

type Episode struct {
	ID            int64     `json:"id"`
	EpisodeNumber int64     `json:"episode_number"`
	EpisodeDate   time.Time `json:"episode_date"`
}
