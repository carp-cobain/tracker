package model

import (
	"strings"

	"github.com/carp-cobain/tracker/domain"
)

// Referral represents a blockchain account that signed up using a referral campaign.
type Referral struct {
	Model
	Campaign   Campaign       `gorm:"foreignKey:CampaignID"`
	CampaignID string         `gorm:"uniqueIndex:campaign_id_account_idx;index;not null"`
	Account    string         `gorm:"uniqueIndex:campaign_id_account_idx;index;not null"`
	Status     ReferralStatus `gorm:"index;not null"`
}

// NewReferral creates a new referral for a campaign.
func NewReferral(campaignID, account string) Referral {
	return Referral{
		CampaignID: campaignID,
		Account:    account,
		Status:     ReferralStatusPending,
	}
}

// ToDomain converts a model to a domain object representation.
func (self Referral) ToDomain() domain.Referral {
	return domain.Referral{
		ID:         domain.MustParseReferralID(self.ID),
		CampaignID: domain.MustParseCampaignID(self.CampaignID),
		Account:    domain.MustValidateAccount(self.Account),
		Status:     self.Status.ToDomain(),
		CreatedAt:  self.CreatedAt.ToDomain(),
		UpdatedAt:  self.UpdatedAt.ToDomain(),
	}
}

// ReferralStatus categorizes referrals
type ReferralStatus int

const (
	// ReferralStatusCanceled means a referral could not be verified or paid (no bonus issued)
	ReferralStatusCanceled ReferralStatus = iota
	// ReferralStatusPending means a referral needs to be verified
	ReferralStatusPending
	// ReferralStatusVerified means a referee has passed kyc and traded crypto
	ReferralStatusVerified
	// ReferralStatusPaid means bonus has been issued for a verified referral
	ReferralStatusPaid
)

// ToDomain converts a referral status to a string.
func (self ReferralStatus) ToDomain() (value domain.ReferralStatus) {
	switch self {
	case ReferralStatusPending:
		value = domain.PendingStatus
	case ReferralStatusVerified:
		value = domain.VerifiedStatus
	case ReferralStatusPaid:
		value = domain.PaidStatus
	case ReferralStatusCanceled:
		value = domain.CanceledStatus
	}
	return
}

// ReferralStatusFromDomain creates a campaign type from a string.
func ReferralStatusFromDomain(value domain.ReferralStatus) (referralStatus ReferralStatus) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case domain.VerifiedStatus:
		referralStatus = ReferralStatusVerified
	case domain.PaidStatus:
		referralStatus = ReferralStatusPaid
	case domain.CanceledStatus:
		referralStatus = ReferralStatusCanceled
	default:
		referralStatus = ReferralStatusPending
	}
	return
}
