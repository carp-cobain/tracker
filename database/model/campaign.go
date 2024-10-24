package model

import "github.com/carp-cobain/tracker/domain"

// Campaign represents a typed campaign for a blockchain account.
type Campaign struct {
	ID        uint64 `gorm:"primarykey"`
	Name      string
	Account   string `gorm:"index;not null"`
	CreatedAt DateTime
	UpdatedAt DateTime
	ExpiresAt DateTime `gorm:"index"`
}

// NewCampaign creates a new referral campaign for a blockchain account.
func NewCampaign(account, name string) Campaign {
	return Campaign{
		Account:   account,
		Name:      name,
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
		CreatedAt: self.CreatedAt.ToDomain(),
		UpdatedAt: self.UpdatedAt.ToDomain(),
		ExpiresAt: self.ExpiresAt.ToDomain(),
	}
}
