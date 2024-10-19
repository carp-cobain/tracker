package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/carp-cobain/tracker/dto"
	"github.com/carp-cobain/tracker/keeper"

	"github.com/gin-gonic/gin"
)

// CampaignHandler is the http/json api for managing campaigns
type CampaignHandler struct {
	campaignKeeper keeper.CampaignKeeper
}

// NewCampaignHandler creates a new campaign handler
func NewCampaignHandler(campaignKeeper keeper.CampaignKeeper) CampaignHandler {
	return CampaignHandler{campaignKeeper}
}

// GET /campaigns
// GetCampaigns gets a page of campaigns for a blockchain account
func (self CampaignHandler) GetCampaigns(c *gin.Context) {
	account := strings.TrimSpace(c.Query("account"))
	if err := dto.ValidateAccount(account); err != nil {
		badRequestJson(c, fmt.Errorf("account query param is missing or invalid"))
		return
	}
	cursor, limit := getPageParams(c)
	next, campaigns := self.campaignKeeper.GetCampaigns(account, cursor, limit)
	okJson(c, gin.H{"cursor": next, "campaigns": campaigns})
}

// GET /campaigns/:id
// GetCampaign gets campaigns by ID
func (self CampaignHandler) GetCampaign(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequestJson(c, err)
		return
	}
	campaign, err := self.campaignKeeper.GetCampaign(id)
	if err != nil {
		notFoundJson(c, err)
		return
	}
	okJson(c, gin.H{"campaign": campaign})
}

// POST /campaigns
// CreateCampaign creates new named campaigns
func (self CampaignHandler) CreateCampaign(c *gin.Context) {
	var request dto.CampaignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequestJson(c, err)
		return
	}
	account, name, err := request.Validate()
	if err != nil {
		badRequestJson(c, err)
		return
	}
	campaign, err := self.campaignKeeper.CreateCampaign(account, name)
	if err != nil {
		badRequestJson(c, err)
		return
	}
	okJson(c, gin.H{"campaign": campaign})
}

// DELETE /campaigns/:id
// ExpireCampaign marks campaigns as expired
func (self CampaignHandler) ExpireCampaign(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequestJson(c, err)
		return
	}
	if err := self.campaignKeeper.ExpireCampaign(id); err != nil {
		badRequestJson(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
