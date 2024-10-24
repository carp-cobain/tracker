package domain

import "time"

// Campaign represents a named campaign for a blockchain account.
type Campaign struct {
	ID        uint64    `json:"id"`
	Account   string    `json:"account"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}
