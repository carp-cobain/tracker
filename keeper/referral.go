package keeper

import (
	"github.com/carp-cobain/tracker/domain"
)

// ReferralPage is a page of referrals
type ReferralPage = domain.Page[domain.Referral]

// ReferralKeeper manages campaign referrals
type ReferralKeeper interface {
	ReferralReader
	ReferralWriter
}

// ReferralReader reads campaign referrals
type ReferralReader interface {
	GetReferrals(campaignID domain.CampaignID, pageParams domain.PageParams) ReferralPage
	GetReferralsWithStatus(status string, pageParams domain.PageParams) ReferralPage
}

// ReferralWriter writes campaign referrals
type ReferralWriter interface {
	CreateReferral(campaignID domain.CampaignID, account domain.Account) (domain.Referral, error)
	UpdateReferral(referralID domain.ReferralID, status string) (domain.Referral, error)
}
