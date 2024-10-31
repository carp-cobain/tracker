package processor

// Helpers for using referral types as strings.
// Copied from model package to not have the low-level dependency in this package
var (
	pendingStatus  = "pending"
	verifiedStatus = "verified"
	canceledStatus = "canceled"
	paidStatus     = "paid"
)
