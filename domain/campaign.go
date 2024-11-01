package domain

import (
	"time"

	"github.com/google/uuid"
)

// Campaign represents a named campaign for a blockchain account.
type Campaign struct {
	ID        CampaignID   `json:"id"`
	Account   Account      `json:"account"`
	Name      string       `json:"name"`
	Type      CampaignType `json:"type"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	ExpiresAt time.Time    `json:"expiresAt"`
}

// CampaignID is a campaign unique identifier.
type CampaignID struct {
	inner uuid.UUID
}

// ParseCampaignID creates a CampaignID from a UUID string.
// This function returns an error if the UUID string cannot be parsed.
func ParseCampaignID(value string) (campaignID CampaignID, err error) {
	var inner uuid.UUID
	inner, err = uuid.Parse(value)
	if err != nil {
		return
	}
	campaignID = CampaignID{inner}
	return
}

// MustParseCampaignID creates a CampaignID from a UUID string.
// This function panics if the UUID string cannot be parsed.
func MustParseCampaignID(value string) CampaignID {
	return CampaignID{uuid.MustParse(value)}
}

// String returns the inner campaign UUID value as a string
func (self CampaignID) String() string {
	return self.inner.String()
}

// MarshalText implements encoding.TextMarshaler.
func (self CampaignID) MarshalText() ([]byte, error) {
	return self.inner.MarshalText()
}
