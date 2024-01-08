package model

type User struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Email         *string `json:"email"`
	EmailVerified *string `json:"emailVerified"`
	Image         *string `json:"image"`
	IsAdmin       bool    `json:"is_admin"`
}

type UserPick struct {
	ID       int64  `json:"id"`
	UserID   string `json:"user_id"`
	PlayerID int64  `json:"player_id"`
}

// For API
type UserWithPicks struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Email         *string `json:"email"`
	EmailVerified *string `json:"emailVerified"`
	Image         *string `json:"image"`
	IsAdmin       bool    `json:"is_admin"`
	PlayerIDs     []int64 `json:"player_ids"`
}
