package domain

import (
	"time"

	"github.com/google/uuid"
)

// Campaign represents a referral account for a campaign.
type Referral struct {
	ID         ReferralID     `json:"id"`
	CampaignID CampaignID     `json:"campaignId"`
	Account    Account        `json:"account"`
	Status     ReferralStatus `json:"status"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}

// ReferralID is a referral unique identifier.
type ReferralID struct {
	inner uuid.UUID
}

// ParseReferralID creates a ReferralID from a UUID string.
// This function returns an error if the UUID string cannot be parsed.
func ParseReferralID(value string) (referralID ReferralID, err error) {
	var inner uuid.UUID
	inner, err = uuid.Parse(value)
	if err != nil {
		return
	}
	referralID = ReferralID{inner}
	return
}

// MustParseReferralID creates a ReferralID from a UUID string.
// This function panics if the UUID string cannot be parsed.
func MustParseReferralID(value string) ReferralID {
	return ReferralID{uuid.MustParse(value)}
}

// String returns the inner referral UUID value as a string
func (self ReferralID) String() string {
	return self.inner.String()
}

// MarshalText implements encoding.TextMarshaler.
func (self ReferralID) MarshalText() ([]byte, error) {
	return self.inner.MarshalText()
}
