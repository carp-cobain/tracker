package model

import (
	"strings"
	"time"
)

// DateTime is used to store timestamps as INT columns in SQLite
type DateTime int64

// ToDomain creates a go stdlib time from the SQLite column type (unix seconds).
func (t *DateTime) ToDomain() time.Time {
	return time.Unix(int64(*t), 0).UTC()
}

// Now returns the current time as a SQLite column type.
func Now() DateTime {
	return DateTime(time.Now().Unix())
}

// Expiry returns one year from now as a SQLite column type.
func Expiry() DateTime {
	return DateTime(time.Now().Add(365 * 24 * time.Hour).Unix())
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
func (self CampaignType) ToDomain() (value string) {
	switch self {
	case CampaignTypeReferral:
		value = "referral"
	case CampaignTypeRewards:
		value = "rewards"
	case CampaignTypeMarketing:
		value = "marketing"
	}
	return
}

// CampaignTypeFromString creates a campaign type from a string.
func CampaignTypeFromString(value string) (campaignType CampaignType) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "rewards":
		campaignType = CampaignTypeRewards
	case "marketing":
		campaignType = CampaignTypeMarketing
	default:
		campaignType = CampaignTypeReferral
	}
	return
}

// ReferralStatus categorizes referrals
type ReferralStatus int

const (
	// ReferralStatusPending means a referral needs to be verified
	ReferralStatusPending ReferralStatus = iota
	// ReferralStatusVerified means a referee has passed kyc and traded crypto
	ReferralStatusVerified
	// ReferralStatusPaid means bonus has been issued for a verified referral
	ReferralStatusPaid
	// ReferralStatusCanceled means a referral could not be verified (no bonus issued)
	ReferralStatusCanceled
)

// ToDomain converts a referral status to a string.
func (self ReferralStatus) ToDomain() (value string) {
	switch self {
	case ReferralStatusPending:
		value = "pending"
	case ReferralStatusVerified:
		value = "verified"
	case ReferralStatusPaid:
		value = "paid"
	case ReferralStatusCanceled:
		value = "canceled"
	}
	return
}

// ReferralStatusFromString creates a campaign type from a string.
func ReferralStatusFromString(value string) (referralStatus ReferralStatus) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "verified":
		referralStatus = ReferralStatusVerified
	case "paid":
		referralStatus = ReferralStatusPaid
	case "canceled":
		referralStatus = ReferralStatusCanceled
	default:
		referralStatus = ReferralStatusPending
	}
	return
}
