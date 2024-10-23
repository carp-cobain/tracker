package keeper

import "github.com/carp-cobain/tracker/domain"

// ReferralKeeper manages campaign referrals
type ReferralKeeper interface {
	ReferralReader
	ReferralWriter
}

// ReferralReader reads campaign referrals
type ReferralReader interface {
	GetReferrals(campaignID uint64, pageParams domain.PageParams) (uint64, []domain.Referral)
	GetReferralsWithStatus(status string, pageParams domain.PageParams) (uint64, []domain.Referral)
}

// ReferralWriter writes campaign referrals
type ReferralWriter interface {
	CreateReferral(campaignID uint64, account string) (domain.Referral, error)
	UpdateReferral(referralID uint64, status string) (domain.Referral, error)
}
