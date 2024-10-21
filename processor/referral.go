package processor

import (
	"log"
	"math/rand"

	"github.com/carp-cobain/tracker/database/model"
	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/keeper"
)

// ReferralVerifier verifies that referrals have placed an order on the exchange.
type ReferralVerifier struct {
	referralKeeper keeper.ReferralKeeper
	cursor         uint64
}

// NewReferralVerifier creates a new referral verifier.
func NewReferralVerifier(referralKeeper keeper.ReferralKeeper) ReferralVerifier {
	return ReferralVerifier{referralKeeper, 0}
}

// VerifyReferrals verifies whether accounts for referrals in a "pending" status have made a trade.
func (self *ReferralVerifier) VerifyReferrals() {
	limit := 1000 // page size
	pending := model.ReferralStatusPending.ToDomain()
	cursor, referrals := self.referralKeeper.GetReferralsWithStatus(pending, self.cursor, limit)
	for _, referral := range referrals {
		self.verifyReferral(referral)
	}
	self.cursor = cursor
}

// verfiy referral logic
func (self *ReferralVerifier) verifyReferral(referral domain.Referral) {
	status := randStatus()
	log.Printf("setting referral %d status to %s", referral.ID, status)
	if _, err := self.referralKeeper.UpdateReferral(referral.ID, status); err != nil {
		log.Printf(
			"failed to update referral %d to status %s: %s",
			referral.ID,
			status,
			err.Error(),
		)
	}
}

// return a pseudo-random referral status
func randStatus() (status string) {
	if rand.Float32() < 0.8 {
		status = model.ReferralStatusVerified.ToDomain()
	} else {
		status = model.ReferralStatusCanceled.ToDomain()
	}
	return
}
