package domain

// CampaignType categorizes campaigns
type CampaignType = string

var (
	// ReferralType means both referer and referee get a bonus
	ReferralType CampaignType = "referral"
	// RewardsType means only the referee gets bonus
	RewardsType CampaignType = "rewards"
	// MarketingType is just a classifier for basic tracking purposes (no bonus paid)
	MarketingType CampaignType = "marketing"
)

// ReferralStatus categorizes referrals
type ReferralStatus = string

var (
	// PendingStatus means a referral needs to be verified
	PendingStatus ReferralStatus = "pending"
	// VerifiedStatus means a referee has passed kyc
	VerifiedStatus ReferralStatus = "verified"
	// PaidStatus means bonus has been issued for a verified referral
	PaidStatus ReferralStatus = "paid"
	// CanceledStatus means a referral could not be verified (no bonus issued)
	CanceledStatus ReferralStatus = "canceled"
)
