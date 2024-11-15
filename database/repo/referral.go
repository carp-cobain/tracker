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

// NewReferralRepoRW creates a new repository for managing referrals for campaigns.
func NewReferralRepoRW(readDB, writeDB *gorm.DB) ReferralRepo {
	return ReferralRepo{readDB, writeDB}
}

// NewReferralRepo creates a new repository for managing referrals for campaigns
// with only a single db pointer.
func NewReferralRepo(db *gorm.DB) ReferralRepo {
	return ReferralRepo{db, db}
}

// GetReferrals gets a page of referrals for a campaign.
func (self ReferralRepo) GetReferrals(
	campaignID domain.CampaignID, pageParams domain.PageParams) domain.Page[domain.Referral] {

	var nextCursor uint64
	results := query.SelectReferrals(self.readDB, campaignID.String(), pageParams.Cursor, pageParams.Limit)
	referrals := make([]domain.Referral, len(results))
	for i, r := range results {
		referrals[i] = r.ToDomain()
		nextCursor = max(nextCursor, uint64(r.CreatedAt))
	}
	return domain.NewPage(nextCursor, pageParams.Limit, referrals)
}

// GetReferralsWithStatus gets a page of referrals with a given status.
func (self ReferralRepo) GetReferralsWithStatus(
	status domain.ReferralStatus, pageParams domain.PageParams) domain.Page[domain.Referral] {

	var nextCursor uint64
	results := query.SelectReferralsWithStatus(self.readDB, status, pageParams.Cursor, pageParams.Limit)
	referrals := make([]domain.Referral, len(results))
	for i, r := range results {
		referrals[i] = r.ToDomain()
		nextCursor = max(nextCursor, uint64(r.CreatedAt))
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
	var result model.Referral
	result, err = query.InsertReferral(self.writeDB, campaignID.String(), account.String())
	if err == nil {
		referral = result.ToDomain()
	}
	return
}

// UpdateReferral updates the status of a referral for a campaign.
func (self ReferralRepo) UpdateReferral(
	referralID domain.ReferralID, status domain.ReferralStatus) (referral domain.Referral, err error) {

	var result model.Referral
	result, err = query.UpdateReferralStatus(self.writeDB, referralID.String(), status)
	if err == nil {
		referral = result.ToDomain()
	}
	return
}
