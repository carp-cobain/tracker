package repo

import (
	"fmt"

	"github.com/carp-cobain/tracker/database/model"
	"github.com/carp-cobain/tracker/database/query"
	"github.com/carp-cobain/tracker/domain"

	"gorm.io/gorm"
)

// ReferralRepo manages referrals for campaigns.
type ReferralRepo struct {
	readDB  *gorm.DB
	writeDB *gorm.DB
}

// NewReferralRepo creates a new repository for managing referrals for campaigns.
func NewReferralRepo(readDB, writeDB *gorm.DB) ReferralRepo {
	return ReferralRepo{readDB, writeDB}
}

// GetReferrals gets a page of referrals for a campaign.
func (self ReferralRepo) GetReferrals(campaignID uint64, pageParams domain.PageParams) ReferralPage {
	var nextCursor uint64
	models := query.SelectReferrals(self.readDB, campaignID, pageParams.Cursor, pageParams.Limit)
	referrals := make([]domain.Referral, len(models))
	for i, model := range models {
		referrals[i] = model.ToDomain()
		nextCursor = max(nextCursor, model.ID)
	}
	return domain.NewPage(nextCursor, pageParams.Limit, referrals)
}

// GetReferralsWithStatus gets a page of referrals with a given status.
func (self ReferralRepo) GetReferralsWithStatus(status string, pageParams domain.PageParams) ReferralPage {
	var nextCursor uint64
	cursor, limit := pageParams.Cursor, pageParams.Limit
	models := query.SelectReferralsWithStatus(self.readDB, status, cursor, limit)
	referrals := make([]domain.Referral, len(models))
	for i, model := range models {
		referrals[i] = model.ToDomain()
		nextCursor = max(nextCursor, model.ID)
	}
	return domain.NewPage(nextCursor, pageParams.Limit, referrals)
}

// CreateReferral creates a referral for a campaign.
func (self ReferralRepo) CreateReferral(campaignID uint64, account string) (referral domain.Referral, err error) {
	var campaign model.Campaign
	campaign, err = query.SelectCampaign(self.readDB, campaignID)
	if err != nil {
		err = fmt.Errorf("campaign %d: %s", campaignID, err.Error())
		return
	}
	if campaign.Type != model.CampaignTypeReferral {
		err = fmt.Errorf("error: non-referral campaign: %s", campaign.Type.ToDomain())
		return
	}
	if campaign.Account == account {
		err = fmt.Errorf("self referral error: %s", account)
		return
	}
	var model model.Referral
	if model, err = query.InsertReferral(self.writeDB, campaignID, account); err == nil {
		referral = model.ToDomain()
	}
	return
}

// UpdateReferral updates the status of a referral for a campaign.
func (self ReferralRepo) UpdateReferral(referralID uint64, statusValue string) (referral domain.Referral, err error) {
	status := model.ReferralStatusFromString(statusValue)
	var model model.Referral
	model, err = query.UpdateReferralStatus(self.writeDB, referralID, status)
	if err == nil {
		referral = model.ToDomain()
	}
	return
}
