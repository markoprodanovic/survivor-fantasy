package model

type SimpleEpisodePoint struct {
	CastID int64 `json:"castId"`
	Points int64 `json:"points"`
}

type EpisodeWithPoints struct {
	ID            int64                `json:"id"`
	EpisodeNumber int64                `json:"episode_number"`
	EpisodeDate   SQLiteDate           `json:"episode_date"`
	Points        []SimpleEpisodePoint `json:"points"`
}
