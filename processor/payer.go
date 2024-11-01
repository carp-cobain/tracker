package processor

import (
	"log"
	"time"

	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/service"
)

// ReferralPayer makes payments for verified referrals.
type ReferralPayer struct {
	referralService service.ReferralService
	batchSize       int
	pageCursor      uint64
}

// NewReferralPayer creates a new referral payment processor.
func NewReferralPayer(
	referralService service.ReferralService,
	batchSize int,
	startCursor uint64,
) ReferralPayer {
	return ReferralPayer{referralService, batchSize, startCursor}
}

// PayVerifiedReferrals makes payments for verified referrals.
func (self ReferralPayer) PayVerifiedReferrals() {
	pageParams := domain.NewPageParams(self.pageCursor, self.batchSize)
	page := self.referralService.GetReferralsWithStatus(domain.VerifiedStatus, pageParams)
	for _, referral := range page.Data {
		self.makeReferralPayment(referral)
	}
	self.pageCursor = page.Cursor

}

// TODO: payment logic would go here
func (self *ReferralPayer) makeReferralPayment(referral domain.Referral) {
	// simulate broadcasting a cosmos blockchain transaction
	log.Printf("simulate blockchain payment to account: %s", referral.Account)
	broadcastTime, _ := time.ParseDuration("5s")
	time.Sleep(broadcastTime)
	// all referrals just get marked as "paid" in this POC
	paid := domain.PaidStatus
	log.Printf("setting referral %s status to %s", referral.ID, paid)
	if _, err := self.referralService.UpdateReferral(referral.ID, paid); err != nil {
		log.Printf(
			"failed to update referral %s to status %s: %s",
			referral.ID,
			paid,
			err.Error(),
		)
	}
}
