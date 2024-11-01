package service

import (
	"github.com/carp-cobain/tracker/domain"
)

// ReferralService manages campaign referrals
type ReferralService interface {
	ReferralReader
	ReferralWriter
}

// ReferralReader reads campaign referrals
type ReferralReader interface {
	GetReferrals(campaignID domain.CampaignID, pageParams domain.PageParams) domain.Page[domain.Referral]
	GetReferralsWithStatus(status domain.ReferralStatus, pageParams domain.PageParams) domain.Page[domain.Referral]
}

// ReferralWriter writes campaign referrals
type ReferralWriter interface {
	CreateReferral(campaignID domain.CampaignID, account domain.Account) (domain.Referral, error)
	UpdateReferral(referralID domain.ReferralID, status domain.ReferralStatus) (domain.Referral, error)
}
