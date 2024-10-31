package processor

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/keeper"
)

// ReferralVerifier verifies that referrals have placed an order on the exchange.
type ReferralVerifier struct {
	referralKeeper keeper.ReferralKeeper
	batchSize      int
	pageCursor     uint64
}

// NewReferralVerifier creates a new referral verifier.
func NewReferralVerifier(
	referralKeeper keeper.ReferralKeeper,
	batchSize int,
	startCursor uint64,
) ReferralVerifier {
	return ReferralVerifier{referralKeeper, batchSize, startCursor}
}

// VerifyReferrals verifies whether accounts for referrals in a "pending" status have passed kyc
// and made a trade.
func (self *ReferralVerifier) VerifyReferrals() {
	pageParams := domain.NewPageParams(self.pageCursor, self.batchSize)
	page := self.referralKeeper.GetReferralsWithStatus(pendingStatus, pageParams)
	for _, referral := range page.Data {
		self.verifyReferral(referral)
	}
	self.pageCursor = page.Cursor
}

// verfiy referral logic
func (self *ReferralVerifier) verifyReferral(referral domain.Referral) {
	status := verifyAccountStatus(referral.Account)
	log.Printf("setting referral %s status to %s", referral.ID, status)
	if _, err := self.referralKeeper.UpdateReferral(referral.ID, status); err != nil {
		log.Printf(
			"failed to update referral %s to status %s: %s",
			referral.ID,
			status,
			err.Error(),
		)
	}
}

// return a pseudo-random referral status
func verifyAccountStatus(account domain.Account) (status string) {
	log.Printf("getting status for account: %s", account)
	// simulate latency
	ms, _ := time.ParseDuration(fmt.Sprintf("%dms", rand.Intn(250)))
	time.Sleep(ms)
	// simulate ~80% success rate
	if rand.Float32() < 0.8 {
		status = verifiedStatus
	} else {
		status = canceledStatus
	}
	return
}
