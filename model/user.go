package model

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserPick struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"user_id"`
	PlayerID int64 `json:"player_id"`
}

// For API
type UserWithPicks struct {
	ID        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	PlayerIDs []int64 `json:"player_ids"`
}
