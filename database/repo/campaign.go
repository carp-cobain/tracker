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
	var result model.Campaign
	if result, err = query.SelectCampaign(self.readDB, id); err == nil {
		campaign = result.ToDomain()
	}
	return
}

// GetCampaigns gets a page of campaigns for a blockchain account
func (self CampaignRepo) GetCampaigns(
	account string,
	pageParams domain.PageParams,
) (
	nextCursor uint64,
	campaigns []domain.Campaign,
) {
	results := query.SelectCampaigns(self.readDB, account, pageParams.Cursor, pageParams.Limit)
	campaigns = make([]domain.Campaign, len(results))
	for i, result := range results {
		campaigns[i] = result.ToDomain()
		nextCursor = max(nextCursor, result.ID)
	}
	return
}

// CreateCampaign creates a new named campaign
func (self CampaignRepo) CreateCampaign(
	account string,
	name string,
) (
	campaign domain.Campaign,
	err error,
) {
	var result model.Campaign
	if result, err = query.InsertCampaign(self.writeDB, account, name); err == nil {
		campaign = result.ToDomain()
	}
	return
}

// UpdateCampaign updates campaign fields.
func (self CampaignRepo) UpdateCampaign(
	id uint64,
	name string,
	expiresAt time.Time,
) (
	campaign domain.Campaign,
	err error,
) {
	var result model.Campaign
	expiry := model.DateTime(expiresAt.Unix())
	if result, err = query.UpdateCampaign(self.writeDB, id, name, expiry); err == nil {
		campaign = result.ToDomain()
	}
	return
}
