package keeper

import "github.com/carp-cobain/tracker/domain"

// CampaignKeeper manages campaigns
type CampaignKeeper interface {
	CampaignReader
	CampaignWriter
}

// CampaignReader reads campaigns
type CampaignReader interface {
	GetCampaign(id uint64) (campaign domain.Campaign, err error)
	GetCampaigns(account string, cursor uint64, limit int) (uint64, []domain.Campaign)
}

// CampaignWriter writes campaigns
type CampaignWriter interface {
	CreateCampaign(account, name string) (domain.Campaign, error)
	ExpireCampaign(id uint64) error
}
