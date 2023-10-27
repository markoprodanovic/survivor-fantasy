package model

type EpisodePoints struct {
	ID        int64 `json:"id"`
	EpisodeID int64 `json:"episode_id"`
	CastID    int64 `json:"cast_id"`
	Points    int64 `json:"points"`
}
