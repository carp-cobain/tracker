package repo

import (
	"time"

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
func (self CampaignRepo) GetCampaigns(account string, pageParams domain.PageParams) CampaignPage {
	var nextCursor uint64
	models := query.SelectCampaigns(self.readDB, account, pageParams.Cursor, pageParams.Limit)
	campaigns := make([]domain.Campaign, len(models))
	for i, model := range models {
		campaigns[i] = model.ToDomain()
		nextCursor = max(nextCursor, model.ID)
	}
	return domain.NewPage(nextCursor, pageParams.Limit, campaigns)
}

// CreateCampaign creates a new named campaign
func (self CampaignRepo) CreateCampaign(account, name string) (campaign domain.Campaign, err error) {
	var model model.Campaign
	if model, err = query.InsertCampaign(self.writeDB, account, name); err == nil {
		campaign = model.ToDomain()
	}
	return
}

// UpdateCampaign updates campaign fields.
func (self CampaignRepo) UpdateCampaign(
	id uint64, name string, expiresAt time.Time) (campaign domain.Campaign, err error) {

	expiry := model.DateTime(expiresAt.Unix())
	var model model.Campaign
	if model, err = query.UpdateCampaign(self.writeDB, id, name, expiry); err == nil {
		campaign = model.ToDomain()
	}
	return
}
