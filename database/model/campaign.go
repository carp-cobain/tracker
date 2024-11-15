package model

import (
	"strings"

	"github.com/carp-cobain/tracker/domain"
)

// Campaign represents a typed campaign for a blockchain account.
type Campaign struct {
	Model
	Name      string
	Account   string       `gorm:"index;not null"`
	Type      CampaignType `gorm:"index"`
	ExpiresAt DateTime     `gorm:"index"`
}

// NewCampaign creates a new referral campaign for a blockchain account.
func NewCampaign(account, name string) Campaign {
	return Campaign{
		Account:   account,
		Name:      name,
		Type:      CampaignTypeReferral,
		ExpiresAt: Expiry(),
	}
}

// ToDomain converts a model to a domain object representation.
func (self Campaign) ToDomain() domain.Campaign {
	return domain.Campaign{
		ID:        domain.MustParseCampaignID(self.ID),
		Account:   domain.MustValidateAccount(self.Account),
		Name:      self.Name,
		Type:      self.Type.ToDomain(),
		CreatedAt: self.CreatedAt.ToDomain(),
		UpdatedAt: self.UpdatedAt.ToDomain(),
		ExpiresAt: self.ExpiresAt.ToDomain(),
	}
}

// IsExpired returns true when a campaign has expired.
func (self Campaign) IsExpired() bool {
	return self.ExpiresAt <= Now()
}

// CampaignType categorizes campaigns
type CampaignType int

const (
	_ CampaignType = iota
	// CampaignTypeReferral means both referer and referee get a bonus
	CampaignTypeReferral
	// CampaignTypeMarketing just a classifier for marketing purposes (no bonus)
	CampaignTypeMarketing
	// CampaignTypeRewards means only the referee gets bonus
	CampaignTypeRewards
)

// ToDomain converts a campaign type to a string.
func (self CampaignType) ToDomain() (value domain.CampaignType) {
	switch self {
	case CampaignTypeReferral:
		value = domain.ReferralType
	case CampaignTypeRewards:
		value = domain.RewardsType
	case CampaignTypeMarketing:
		value = domain.MarketingType
	}
	return
}

// CampaignTypeFromDomain creates a campaign type from a string.
func CampaignTypeFromDomain(value domain.CampaignType) (campaignType CampaignType) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case domain.RewardsType:
		campaignType = CampaignTypeRewards
	case domain.MarketingType:
		campaignType = CampaignTypeMarketing
	default:
		campaignType = CampaignTypeReferral
	}
	return
}
