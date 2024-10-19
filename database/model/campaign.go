package model

import "github.com/carp-cobain/tracker/domain"

// Campaign represents a typed campaign for a blockchain account.
type Campaign struct {
	ID        uint64 `gorm:"primarykey"`
	Name      string
	Account   string       `gorm:"index;not null"`
	Type      CampaignType `gorm:"index;default 0"`
	CreatedAt DateTime
	UpdatedAt DateTime
	ExpiresAt DateTime `gorm:"index"`
}

// NewCampaign creates a new referral campaign for a blockchain account.
func NewCampaign(account, name string) Campaign {
	return NewCampaignWithType(account, name, CampaignTypeReferral)
}

// NewCampaignWithType creates a new campaign with a given type for a blockchain account.
func NewCampaignWithType(account, name string, campaignType CampaignType) Campaign {
	return Campaign{
		Account:   account,
		Name:      name,
		Type:      campaignType,
		CreatedAt: Now(),
		UpdatedAt: Now(),
		ExpiresAt: Expiry(),
	}
}

// ToDomain converts a model to a domain object representation.
func (self Campaign) ToDomain() domain.Campaign {
	return domain.Campaign{
		ID:        self.ID,
		Account:   self.Account,
		Name:      self.Name,
		Type:      self.Type.ToDomain(),
		CreatedAt: self.CreatedAt.ToDomain(),
		UpdatedAt: self.UpdatedAt.ToDomain(),
		ExpiresAt: self.ExpiresAt.ToDomain(),
	}
}
