package repo

import (
	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/keeper"
)

// Type checks: ensure repos implement keeper interfaces
var _ keeper.CampaignKeeper = &CampaignRepo{}
var _ keeper.ReferralKeeper = &ReferralRepo{}

// CampaignPage is a type alias for `domain.Page[domain.Campaign]`
type CampaignPage = domain.Page[domain.Campaign]

// ReferralPage is a type alias for `domain.Page[domain.Referral]`
type ReferralPage = domain.Page[domain.Referral]
