package processor

import (
	"log"
	"time"

	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/keeper"
)

// ReferralPayer makes payments for verified referrals.
type ReferralPayer struct {
	referralKeeper keeper.ReferralKeeper
	batchSize      int
	pageCursor     uint64
}

// NewReferralPayer creates a new referral payment processor.
func NewReferralPayer(
	referralKeeper keeper.ReferralKeeper,
	batchSize int,
	startCursor uint64,
) ReferralPayer {
	return ReferralPayer{referralKeeper, batchSize, startCursor}
}

// PayVerifiedReferrals makes payments for verified referrals.
func (self ReferralPayer) PayVerifiedReferrals() {
	log.Println("PayVerifiedReferrals")
	pageParams := domain.NewPageParams(self.pageCursor, self.batchSize)
	page := self.referralKeeper.GetReferralsWithStatus(verifiedStatus, pageParams)
	for _, referral := range page.Data {
		self.makeReferralPayment(referral)
	}
	self.pageCursor = page.Cursor

}

// TODO: payment logic would go here
func (self *ReferralPayer) makeReferralPayment(referral domain.Referral) {
	// simulate broadcasting a cosmos blockchain transaction
	broadcastTime, _ := time.ParseDuration("5s")
	time.Sleep(broadcastTime)
	// all referrals just get marked as processed in this POC
	log.Printf("setting referral %d status to %s", referral.ID, processedStatus)
	if _, err := self.referralKeeper.UpdateReferral(referral.ID, processedStatus); err != nil {
		log.Printf(
			"failed to update referral %d to status %s: %s",
			referral.ID,
			processedStatus,
			err.Error(),
		)
	}
}
