package repo

import (
	"fmt"

	"github.com/carp-cobain/tracker/database/model"
	"github.com/carp-cobain/tracker/database/query"
	"github.com/carp-cobain/tracker/domain"
	"gorm.io/gorm"
)

// CampaignRepo manages campaigns in a database.
type CampaignRepo struct {
	readDB  *gorm.DB
	writeDB *gorm.DB
}

// NewCampaignRepo creates a new repository for managing campaigns.
func NewCampaignRepo(readDB, writeDB *gorm.DB) CampaignRepo {
	return CampaignRepo{readDB, writeDB}
}

// GetCampaign gets a campaign by ID
func (self CampaignRepo) GetCampaign(id uint64) (campaign domain.Campaign, err error) {
	var model model.Campaign
	if model, err = query.SelectCampaign(self.readDB, id); err == nil {
		campaign = model.ToDomain()
	}
	return
}

// GetCampaigns gets a page of campaigns for a blockchain account
func (self CampaignRepo) GetCampaigns(account string, cursor uint64, limit int) (next uint64, campaigns []domain.Campaign) {
	models := query.SelectCampaigns(self.readDB, account, cursor, limit)
	campaigns = make([]domain.Campaign, len(models))
	for i, model := range models {
		campaigns[i] = model.ToDomain()
		next = max(next, model.ID)
	}
	return
}

// CreateCampaign creates a new named campaign
func (self CampaignRepo) CreateCampaign(account, name string) (campaign domain.Campaign, err error) {
	if model, err := query.InsertCampaign(self.writeDB, account, name); err == nil {
		campaign = model.ToDomain()
	} else {
		err = fmt.Errorf("campaign: %s", err.Error())
	}
	return
}

// ExpireCampaign marks a campaign as expired.
func (self CampaignRepo) ExpireCampaign(id uint64) error {
	return query.ExpireCampaign(self.writeDB, id)
}
