package model

import "time"

type URL struct {
	ShortURL  string
	LongURL   string
	ExpiresAt *time.Time
	CreatedAt *time.Time
}
