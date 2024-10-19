package model

import "github.com/carp-cobain/tracker/domain"

// Referral represents a blockchain account that signed up using a referral campaign.
type Referral struct {
	ID         uint64         `gorm:"primarykey"`
	Campaign   Campaign       `gorm:"foreignKey:CampaignID"`
	CampaignID uint64         `gorm:"uniqueIndex:compositeindex;index;not null"`
	Account    string         `gorm:"uniqueIndex:compositeindex;index;not null"`
	Status     ReferralStatus `gorm:"index;not null"`
	CreatedAt  DateTime
	UpdatedAt  DateTime
}

// NewReferral creates a new referral for a campaign.
func NewReferral(campaignID uint64, account string) Referral {
	return Referral{
		CampaignID: campaignID,
		Account:    account,
		Status:     ReferralStatusPending,
		CreatedAt:  Now(),
		UpdatedAt:  Now(),
	}
}

// ToDomain converts a model to a domain object representation.
func (self Referral) ToDomain() domain.Referral {
	return domain.Referral{
		ID:         self.ID,
		CampaignID: self.CampaignID,
		Account:    self.Account,
		Status:     self.Status.ToDomain(),
		CreatedAt:  self.CreatedAt.ToDomain(),
		UpdatedAt:  self.UpdatedAt.ToDomain(),
	}
}
