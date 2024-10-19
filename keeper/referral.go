package keeper

import "github.com/carp-cobain/tracker/domain"

// ReferralKeeper manages campaign referrals
type ReferralKeeper interface {
	ReferralReader
	ReferralWriter
}

// ReferralReader reads campaign referrals
type ReferralReader interface {
	GetReferrals(campaignID, cursor uint64, limit int) (uint64, []domain.Referral)
	GetReferralsWithStatus(status string, cursor uint64, limit int) (uint64, []domain.Referral)
}

// ReferralWriter writes campaign referrals
type ReferralWriter interface {
	CreateReferral(campaignID uint64, account string) (domain.Referral, error)
	SetReferralStatus(referralID uint64, status string) (domain.Referral, error)
}
