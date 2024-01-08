package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type SQLiteDate time.Time

func (sd *SQLiteDate) Scan(value interface{}) error {
	strVal, ok := value.(string)
	if !ok {
		return fmt.Errorf("SQLiteDate must be a string")
	}
	parsedTime, err := time.Parse("2006-01-02", strVal)
	if err != nil {
		return err
	}
	*sd = SQLiteDate(parsedTime)
	return nil
}

func (sd SQLiteDate) Value() (driver.Value, error) {
	t := time.Time(sd)
	return t.Format("2006-01-02"), nil
}

// MarshalJSON will be called when the struct is being serialized to JSON.
func (sd SQLiteDate) MarshalJSON() ([]byte, error) {
	// Format the time as a JSON string.
	t := time.Time(sd)
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(stamp), nil
}

// UnmarshalJSON will be called when the struct is being deserialized from JSON.
func (sd *SQLiteDate) UnmarshalJSON(data []byte) error {
	// Parse the time from the JSON string.
	t, err := time.Parse("\"2006-01-02\"", string(data))
	if err != nil {
		return err
	}
	*sd = SQLiteDate(t)
	return nil
}

type Episode struct {
	ID            int64      `json:"id"`
	EpisodeNumber int64      `json:"episode_number"`
	EpisodeDate   SQLiteDate `json:"episode_date"`
}
