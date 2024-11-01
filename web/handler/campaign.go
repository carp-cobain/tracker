package handler

import (
	"net/http"
	"time"

	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/service"
	"github.com/carp-cobain/tracker/web/dto"

	"github.com/gin-gonic/gin"
)

// CampaignHandler is the http/json api for managing campaigns
type CampaignHandler struct {
	campaignService service.CampaignService
}

// NewCampaignHandler creates a new campaign handler
func NewCampaignHandler(campaignService service.CampaignService) CampaignHandler {
	return CampaignHandler{campaignService}
}

// GET /campaigns
// GetCampaigns gets a page of campaigns for a blockchain account
func (self CampaignHandler) GetCampaigns(c *gin.Context) {
	account, err := domain.NewAccount(c.Query("account")).Validate()
	if err != nil {
		badRequestJson(c, err)
		return
	}
	pageParms := getPageParams(c)
	campaigns := self.campaignService.GetCampaigns(account, pageParms)
	okJson(c, gin.H{"campaigns": campaigns})
}

// GET /campaigns/:id
// GetCampaign gets campaigns by ID
func (self CampaignHandler) GetCampaign(c *gin.Context) {
	campaignID, err := domain.ParseCampaignID(c.Param("id"))
	if err != nil {
		badRequestJson(c, err)
		return
	}
	campaign, err := self.campaignService.GetCampaign(campaignID)
	if err != nil {
		notFoundJson(c, err)
		return
	}
	okJson(c, gin.H{"campaign": campaign})
}

// POST /campaigns
// CreateCampaign creates new named campaigns
func (self CampaignHandler) CreateCampaign(c *gin.Context) {
	var request dto.CreateCampaignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequestJson(c, err)
		return
	}
	account, name, err := request.Validate()
	if err != nil {
		badRequestJson(c, err)
		return
	}
	campaign, err := self.campaignService.CreateCampaign(account, name)
	if err != nil {
		badRequestJson(c, err)
		return
	}
	createdJson(c, gin.H{"campaign": campaign})
}

// DELETE /campaigns/:id
// ExpireCampaign marks campaigns as expired
func (self CampaignHandler) ExpireCampaign(c *gin.Context) {
	campaignID, err := domain.ParseCampaignID(c.Param("id"))
	if err != nil {
		badRequestJson(c, err)
		return
	}
	expiresAt := time.Now()
	if _, err := self.campaignService.UpdateCampaign(campaignID, "", expiresAt); err != nil {
		badRequestJson(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// PATCH /campaigns/:id
// UpdateCampaign updates campaign name and/or expiration.
func (self CampaignHandler) UpdateCampaign(c *gin.Context) {
	campaignID, err := domain.ParseCampaignID(c.Param("id"))
	if err != nil {
		badRequestJson(c, err)
		return
	}
	var request dto.UpdateCampaignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequestJson(c, err)
		return
	}
	name, expiresAt, err := request.Validate()
	if err != nil {
		badRequestJson(c, err)
		return
	}
	campaign, err := self.campaignService.UpdateCampaign(campaignID, name, expiresAt)
	if err != nil {
		badRequestJson(c, err)
		return
	}
	okJson(c, gin.H{"campaign": campaign})
}
