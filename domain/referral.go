package domain

import "time"

type Referral struct {
	ID         uint64    `json:"id"`
	CampaignID uint64    `json:"campaignId"`
	Account    string    `json:"account"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
