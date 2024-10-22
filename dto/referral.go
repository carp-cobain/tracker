package dto

// ReferralRequest is the request type for adding campaign referrals.
type ReferralRequest struct {
	Account string `json:"account" binding:"required"`
}

// Validate referral request fields
func (self ReferralRequest) Validate() (string, error) {
	return ValidateAccount(self.Account)
}
