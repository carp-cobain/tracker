package dto

import (
	"fmt"
	"strings"
)

// CampaignRequest is the request type for creating campaigns.
type CampaignRequest struct {
	Account string `json:"account" binding:"required,min=41,max=61"`
	Name    string `json:"name"`
}

// Validate campaign request address
func (self CampaignRequest) Validate() (string, string, error) {
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
