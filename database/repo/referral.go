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
	campaignID uint64,
	pageParams domain.PageParams,
) (
	nextCursor uint64,
	referrals []domain.Referral,
) {
	results := query.SelectReferrals(self.readDB, campaignID, pageParams.Cursor, pageParams.Limit)
	referrals = make([]domain.Referral, len(results))
	for i, result := range results {
		referrals[i] = result.ToDomain()
		nextCursor = max(nextCursor, result.ID)
	}
	return
}

// GetReferralsWithStatus gets a page of referrals with a given status.
func (self ReferralRepo) GetReferralsWithStatus(
	status string,
	pageParams domain.PageParams,
) (
	nextCursor uint64,
	referrals []domain.Referral,
) {
	cursor, limit := pageParams.Cursor, pageParams.Limit
	results := query.SelectReferralsWithStatus(self.readDB, status, cursor, limit)
	referrals = make([]domain.Referral, len(results))
	for i, result := range results {
		referrals[i] = result.ToDomain()
		nextCursor = max(nextCursor, result.ID)
	}
	return
}

// CreateReferral creates a referral for a campaign.
func (self ReferralRepo) CreateReferral(
	campaignID uint64,
	account string,
) (
	referral domain.Referral,
	err error,
) {
	var campaign model.Campaign
	campaign, err = query.SelectCampaign(self.readDB, campaignID)
	if err != nil {
		err = fmt.Errorf("campaign %d: %s", campaignID, err.Error())
		return
	}
	if campaign.Account == account {
		err = fmt.Errorf("self referral error: %s", account)
		return
	}
	var result model.Referral
	if result, err = query.InsertReferral(self.writeDB, campaignID, account); err == nil {
		referral = result.ToDomain()
	}
	return
}

// UpdateReferral updates the status of a referral for a campaign.
func (self ReferralRepo) UpdateReferral(
	referralID uint64,
	statusValue string,
) (
	referral domain.Referral,
	err error,
) {
	var result model.Referral
	status := model.ReferralStatusFromString(statusValue)
	if result, err = query.UpdateReferralStatus(self.writeDB, referralID, status); err == nil {
		referral = result.ToDomain()
	}
	return
}
