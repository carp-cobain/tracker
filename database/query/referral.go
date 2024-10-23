package query

import (
	"github.com/carp-cobain/tracker/database/model"
	"gorm.io/gorm"
)

// SelectReferral selects a referral by id
func SelectReferral(db *gorm.DB, id uint64) (referral model.Referral, err error) {
	err = db.Where("id = ?", id).First(&referral).Error
	return
}

// SelectReferrals selects all referrals for a campaign.
func SelectReferrals(db *gorm.DB, campaignID, cursor uint64, limit int) (referrals []model.Referral) {
	db.Where("campaign_id = ?", campaignID).
		Where("id > ?", cursor).
		Order("id").
		Limit(limit).
		Find(&referrals)
	return
}

// SelectReferralsWithStatus selects a page of referrals with a given status.
func SelectReferralsWithStatus(db *gorm.DB, status string, cursor uint64, limit int) (referrals []model.Referral) {
	db.Where("status = ?", model.ReferralStatusFromString(status)).
		Where("id > ?", cursor).
		Order("id").
		Limit(limit).
		Find(&referrals)
	return
}

// InsertReferral inserts a new referral for a campaign.
func InsertReferral(db *gorm.DB, campaignID uint64, account string) (referral model.Referral, err error) {
	referral = model.NewReferral(campaignID, account)
	err = db.Create(&referral).Error
	return
}

// UpdateReferralStatus updates referral status.
func UpdateReferralStatus(
	db *gorm.DB, referralID uint64, status model.ReferralStatus) (referral model.Referral, err error) {

	if referral, err = SelectReferral(db, referralID); err == nil {
		data := updates{"status": status}
		err = db.Model(&referral).Updates(data).Error
	}
	return
}
