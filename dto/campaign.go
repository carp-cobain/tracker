package dto

import (
	"fmt"
	"strings"
	"time"
)

// CreateCampaignRequest is the request type for creating campaigns.
type CreateCampaignRequest struct {
	Account string `json:"account" binding:"required,min=41,max=61"`
	Name    string `json:"name"`
}

// Validate campaign request address
func (self CreateCampaignRequest) Validate() (string, string, error) {
	account := strings.TrimSpace(self.Account)
	if err := ValidateAccount(account); err != nil {
		return "", "", err
	}
	name := strings.TrimSpace(self.Name)
	if len(name) > 100 {
		return "", "", fmt.Errorf("campaign name too long: %d > 100", len(name))
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
	expiresAt := self.ExpiresAt
	if expiresAt.Before(time.Now()) {
		expiresAt = time.Unix(0, 0)
	}
	return name, expiresAt, nil
}
