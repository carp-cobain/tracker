package dto

import (
	"fmt"
	"strings"
	"time"

	"github.com/carp-cobain/tracker/domain"
)

// CreateCampaignRequest is the request type for creating campaigns.
type CreateCampaignRequest struct {
	Account domain.Account `json:"account" binding:"required"`
	Name    string         `json:"name"`
}

// Validate campaign request address
func (self CreateCampaignRequest) Validate() (domain.Account, string, error) {
	account := self.Account
	name := strings.TrimSpace(self.Name)
	if len(name) > 100 {
		return account, "", fmt.Errorf("campaign name too long: %d > 100", len(name))
	}
	return account, name, nil
}

// UpdateCampaignRequest is the request type for updating campaigns.
type UpdateCampaignRequest struct {
	Name      string    `json:"name"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// Validate campaign request address
func (self UpdateCampaignRequest) Validate() (string, time.Time, error) {
	name := strings.TrimSpace(self.Name)
	if len(name) > 100 {
		return "", self.ExpiresAt, fmt.Errorf("campaign name too long: %d > 100", len(name))
	}
	if name == "" && self.ExpiresAt.IsZero() {
		return "", self.ExpiresAt, fmt.Errorf("no campaign update provided")
	}
	return name, self.ExpiresAt, nil
}
