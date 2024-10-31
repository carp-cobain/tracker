package keeper

import (
	"time"

	"github.com/carp-cobain/tracker/domain"
)

// CampaignKeeper manages campaigns
type CampaignKeeper interface {
	CampaignReader
	CampaignWriter
}

// CampaignReader reads campaigns
type CampaignReader interface {
	GetCampaign(domain.CampaignID) (domain.Campaign, error)
	GetCampaigns(account domain.Account, pageParams domain.PageParams) domain.Page[domain.Campaign]
}

// CampaignWriter writes campaigns
type CampaignWriter interface {
	CreateCampaign(account domain.Account, name string) (domain.Campaign, error)
	UpdateCampaign(campaignID domain.CampaignID, name string, expiresAt time.Time) (domain.Campaign, error)
}
