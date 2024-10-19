package query

import (
	"github.com/carp-cobain/tracker/database/model"
	"gorm.io/gorm"
)

// SelectCampaign selects a campaign by id
func SelectCampaign(db *gorm.DB, id uint64) (campaign model.Campaign, err error) {
	if err = db.Where("id = ?", id).First(&campaign).Error; err == nil {
		if campaign.ExpiresAt <= model.Now() {
			err = ErrCampaignExpired
		}
	}
	return
}

// SelectCampaigns selects a page of active (ie not expired) campaigns for a blockchain account.
func SelectCampaigns(db *gorm.DB, account string, cursor uint64, limit int) (campaigns []model.Campaign) {
	db.Where("account = ?", account).
		Where("expires_at > ?", model.Now()).
		Where("id > ?", cursor).
		Order("id").
		Limit(limit).
		Find(&campaigns)
	return
}

// InsertCampaign inserts a new named campaign for a blockchain account.
func InsertCampaign(db *gorm.DB, account, name string) (campaign model.Campaign, err error) {
	campaign = model.NewCampaign(account, name)
	err = db.Create(&campaign).Error
	return
}
