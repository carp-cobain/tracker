package handler

import (
	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/service"
	"github.com/carp-cobain/tracker/web/dto"

	"github.com/gin-gonic/gin"
)

// ReferralHandler is the http/json api for managing campaign referrals
type ReferralHandler struct {
	campaignReader  service.CampaignReader
	referralService service.ReferralService
}

// NewReferralHandler creates a new referral campaign handler
func NewReferralHandler(
	campaignReader service.CampaignReader,
	referralService service.ReferralService,
) ReferralHandler {
	return ReferralHandler{campaignReader, referralService}
}

// GET /campaigns/:id/referrals
// GetReferrals gets a page of referrals for a campaign
func (self ReferralHandler) GetReferrals(c *gin.Context) {
	campaignID, err := domain.ParseCampaignID(c.Param("id"))
	if err != nil {
		badRequestJson(c, err)
		return
	}
	pageParams := getPageParams(c)
	referrals := self.referralService.GetReferrals(campaignID, pageParams)
	if referrals.IsEmpty() {
		// Only verify campaign exists when no referrals are found
		if _, err := self.campaignReader.GetCampaign(campaignID); err != nil {
			notFoundJson(c, err)
			return
		}
	}
	okJson(c, gin.H{"referrals": referrals})
}

// POST /campaigns/:id/referrals
// CreateSignup creates a referral for a campaign
func (self ReferralHandler) CreateReferral(c *gin.Context) {
	campaignID, err := domain.ParseCampaignID(c.Param("id"))
	if err != nil {
		badRequestJson(c, err)
		return
	}
	var request dto.ReferralRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequestJson(c, err)
		return
	}
	if _, err := self.campaignReader.GetCampaign(campaignID); err != nil {
		notFoundJson(c, err)
		return
	}
	referral, err := self.referralService.CreateReferral(campaignID, request.Account)
	if err != nil {
		badRequestJson(c, err)
		return
	}
	okJson(c, gin.H{"referral": referral})
}
