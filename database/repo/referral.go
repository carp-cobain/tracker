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
func (self ReferralRepo) GetReferrals(
	campaignID domain.CampaignID, pageParams domain.PageParams) domain.Page[domain.Referral] {

	var nextCursor uint64
	models := query.SelectReferrals(self.readDB, campaignID.String(), pageParams.Cursor, pageParams.Limit)
	referrals := make([]domain.Referral, len(models))
	for i, model := range models {
		referrals[i] = model.ToDomain()
		nextCursor = max(nextCursor, uint64(model.CreatedAt))
	}
	return domain.NewPage(nextCursor, pageParams.Limit, referrals)
}

// GetReferralsWithStatus gets a page of referrals with a given status.
func (self ReferralRepo) GetReferralsWithStatus(
	status string, pageParams domain.PageParams) domain.Page[domain.Referral] {

	var nextCursor uint64
	cursor, limit := pageParams.Cursor, pageParams.Limit
	models := query.SelectReferralsWithStatus(self.readDB, status, cursor, limit)
	referrals := make([]domain.Referral, len(models))
	for i, model := range models {
		referrals[i] = model.ToDomain()
		nextCursor = max(nextCursor, uint64(model.CreatedAt))
	}
	return domain.NewPage(nextCursor, pageParams.Limit, referrals)
}

// CreateReferral creates a referral for a campaign.
func (self ReferralRepo) CreateReferral(
	campaignID domain.CampaignID, account domain.Account) (referral domain.Referral, err error) {

	var campaign model.Campaign
	campaign, err = query.SelectCampaign(self.readDB, campaignID.String())
	if err != nil {
		err = fmt.Errorf("campaign %s: %s", campaignID, err.Error())
		return
	}
	if campaign.Type != model.CampaignTypeReferral {
		err = fmt.Errorf("error: non-referral campaign: %s", campaign.Type.ToDomain())
		return
	}
	if campaign.Account == account.String() {
		err = fmt.Errorf("self referral error: %s", account)
		return
	}
	var model model.Referral
	model, err = query.InsertReferral(self.writeDB, campaignID.String(), account.String())
	if err == nil {
		referral = model.ToDomain()
	}
	return
}

// UpdateReferral updates the status of a referral for a campaign.
func (self ReferralRepo) UpdateReferral(
	referralID domain.ReferralID, status string) (referral domain.Referral, err error) {

	var model model.Referral
	model, err = query.UpdateReferralStatus(self.writeDB, referralID.String(), status)
	if err == nil {
		referral = model.ToDomain()
	}
	return
}
