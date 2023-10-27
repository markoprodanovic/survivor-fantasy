package model

import "time"

type SimpleEpisodePoint struct {
	CastID int `json:"castId"`
	Points int `json:"points"`
}

type EpisodeWithPoints struct {
	EpisodeNumber int64                `json:"episode_number"`
	EpisodeDate   time.Time            `json:"episode_date"`
	Points        []SimpleEpisodePoint `json:"points"`
}
