package dto

import "strings"

// ReferralRequest is the request type for adding campaign referrals.
type ReferralRequest struct {
	Account string `json:"account" binding:"required,min=41,max=61"`
}

// Validate referral request fields
func (self ReferralRequest) Validate() (string, error) {
	account := strings.TrimSpace(self.Account)
	if err := ValidateAccount(account); err != nil {
		return "", err
	}
	return account, nil
}
