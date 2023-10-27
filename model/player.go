package model

type Player struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int64  `json:"age"`
	TribeID    int64  `json:"tribe_id"`
	Eliminated bool   `json:"eliminated"`
}
