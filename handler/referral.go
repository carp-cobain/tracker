package handler

import (
	"github.com/carp-cobain/tracker/dto"
	"github.com/carp-cobain/tracker/keeper"
	"github.com/gin-gonic/gin"
)

// ReferralHandler is the http/json api for managing campaign referrals
type ReferralHandler struct {
	campaignReader keeper.CampaignReader
	referralKeeper keeper.ReferralKeeper
}

// NewReferralHandler creates a new referral campaign handler
func NewReferralHandler(
	campaignReader keeper.CampaignReader,
	referralKeeper keeper.ReferralKeeper,
) ReferralHandler {
	return ReferralHandler{campaignReader, referralKeeper}
}

// GET /campaigns/:id/referrals
// GetReferrals gets a page of referrals for a campaign
func (self ReferralHandler) GetReferrals(c *gin.Context) {
	campaignID, err := uintParam(c, "id")
	if err != nil {
		badRequestJson(c, err)
		return
	}
	nextCursor, referrals := self.referralKeeper.GetReferrals(campaignID, getPageParams(c))
	if len(referrals) == 0 {
		if _, err := self.campaignReader.GetCampaign(campaignID); err != nil {
			notFoundJson(c, err)
			return
		}
	}
	okJson(c, gin.H{"cursor": nextCursor, "referrals": referrals})
}

// POST /campaigns/:id/referrals
// CreateSignup creates a referral for a campaign
func (self ReferralHandler) CreateReferral(c *gin.Context) {
	campaignID, err := uintParam(c, "id")
	if err != nil {
		badRequestJson(c, err)
		return
	}
	var request dto.ReferralRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequestJson(c, err)
		return
	}
	account, err := request.Validate()
	if err != nil {
		badRequestJson(c, err)
		return
	}
	if _, err := self.campaignReader.GetCampaign(campaignID); err != nil {
		notFoundJson(c, err)
		return
	}
	referral, err := self.referralKeeper.CreateReferral(campaignID, account)
	if err != nil {
		badRequestJson(c, err)
		return
	}
	okJson(c, gin.H{"referral": referral})
}
