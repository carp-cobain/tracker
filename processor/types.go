package processor

import "github.com/carp-cobain/tracker/database/model"

// Helpers for using referral types as strings.
var (
	pendingStatus   = model.ReferralStatusPending.ToDomain()
	verifiedStatus  = model.ReferralStatusVerified.ToDomain()
	canceledStatus  = model.ReferralStatusCanceled.ToDomain()
	processedStatus = model.ReferralStatusProcessed.ToDomain()
)
