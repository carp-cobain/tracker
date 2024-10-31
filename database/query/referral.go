package query

import (
	"github.com/carp-cobain/tracker/database/model"
	"gorm.io/gorm"
)

// SelectReferral selects a referral by id
func SelectReferral(db *gorm.DB, referralID string) (referral model.Referral, err error) {
	err = db.Where("id = ?", referralID).First(&referral).Error
	return
}

// SelectReferrals selects all referrals for a campaign.
func SelectReferrals(
	db *gorm.DB, campaignID string, cursor uint64, limit int) (referrals []model.Referral) {

	db.Where("campaign_id = ?", campaignID).
		Where("created_at > ?", cursor).
		Order("created_at").
		Limit(limit).
		Find(&referrals)

	return
}

// SelectReferralsWithStatus selects a page of referrals with a given status.
func SelectReferralsWithStatus(
	db *gorm.DB, status string, cursor uint64, limit int) (referrals []model.Referral) {

	db.Where("status = ?", model.ReferralStatusFromString(status)).
		Where("created_at > ?", cursor).
		Order("created_at").
		Limit(limit).
		Find(&referrals)

	return
}

// InsertReferral inserts a new referral for a campaign.
func InsertReferral(
	db *gorm.DB, campaignID, account string) (referral model.Referral, err error) {

	referral = model.NewReferral(campaignID, account)
	err = db.Create(&referral).Error
	return
}

// UpdateReferralStatus updates referral status.
func UpdateReferralStatus(
	db *gorm.DB,
	referralID string,
	statusValue string,
) (
	referral model.Referral,
	err error,
) {
	status := model.ReferralStatusFromString(statusValue)
	if referral, err = SelectReferral(db, referralID); err == nil {
		err = db.Model(&referral).
			Updates(Changeset{"status": status}).
			Error
	}
	return
}
