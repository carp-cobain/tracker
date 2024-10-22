package processor

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/carp-cobain/tracker/database/model"
	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/keeper"
)

// processor page size
const maxReferrals = 1000

// pending status
var pendingStatus = model.ReferralStatusPending.ToDomain()

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
	nextCursor, referrals :=
		self.referralKeeper.GetReferralsWithStatus(pendingStatus, self.cursor, maxReferrals)
	for _, referral := range referrals {
		self.verifyReferral(referral)
	}
	self.cursor = nextCursor
}

// verfiy referral logic
func (self *ReferralVerifier) verifyReferral(referral domain.Referral) {
	status := getTradeStatus(referral.Account)
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
func getTradeStatus(account string) (status string) {
	log.Printf("getting status for account: %s", account)
	// simulate latency
	ms, _ := time.ParseDuration(fmt.Sprintf("%dms", rand.Intn(200)))
	time.Sleep(ms)
	// simulate ~80% success rate
	if rand.Float32() < 0.8 {
		status = model.ReferralStatusVerified.ToDomain()
	} else {
		status = model.ReferralStatusCanceled.ToDomain()
	}
	return
}
