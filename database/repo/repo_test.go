package repo_test

import (
	"testing"
	"time"

	"github.com/carp-cobain/tracker/database"
	"github.com/carp-cobain/tracker/database/repo"
	"github.com/carp-cobain/tracker/domain"

	"gorm.io/gorm"
)

func createTestDB(t *testing.T) *gorm.DB {
	db, err := database.Connect("file::memory:?cache=shared", 1)
	if err != nil {
		t.Fatalf("failed to connect to test database: %+v", err)
	}
	if err := database.RunMigrations(db); err != nil {
		t.Fatalf("failed to run database migrations: %+v", err)
	}
	return db
}

func TestCampaignRepo(t *testing.T) {
	// Setup
	db := createTestDB(t)
	account := domain.MustValidateAccount("tp1mrzpjszjs6dc5e8fwy23trnz775rwqvhpzzzz2")
	campaignRepo := repo.NewCampaignRepo(db)

	// Create
	campaign, err := campaignRepo.CreateCampaign(account, "Campaign Unit Testing")
	if err != nil {
		t.Fatalf("failed to create campaign: %+v", err)
	}
	if campaign.Type != domain.ReferralType {
		t.Fatalf("expected campaign type: referral, got: %s", campaign.Type)
	}

	// Update
	updatedName := "My Updated Campaign"
	expiresAt := campaign.ExpiresAt.Add(30 * 24 * time.Hour)
	if _, err := campaignRepo.UpdateCampaign(campaign.ID, updatedName, expiresAt); err != nil {
		t.Fatalf("failed to update campaign: %+v", err)
	}

	// Read
	if _, err := campaignRepo.GetCampaign(campaign.ID); err != nil {
		t.Fatalf("failed to get campaign: %+v", err)
	}
	if page := campaignRepo.GetCampaigns(account, domain.DefaultPageParams()); page.Size != 1 {
		t.Fatalf("got unexpected number of campaigns")
	}
}

func TestReferralRepo(t *testing.T) {
	// Setup
	db := createTestDB(t)
	referralRepo := repo.NewReferralRepo(db)

	// Base campaign
	referer := domain.MustValidateAccount("tp1mrzpjszjs6dc5e8fwy23trnz775rwqvhpzzzz3")
	campaign, err := repo.NewCampaignRepo(db).CreateCampaign(referer, "Referral Unit Testing")
	if err != nil {
		t.Fatalf("failed to create campaign: %+v", err)
	}

	// Create
	referee := domain.MustValidateAccount("tp1mrzpjszjs6dc5e8fwy23trnz775rwqvhpzzzz4")
	referral, err := referralRepo.CreateReferral(campaign.ID, referee)
	if err != nil {
		t.Fatalf("failed to create referral: %+v", err)
	}

	// Read
	pageParams := domain.DefaultPageParams()
	if page := referralRepo.GetReferrals(campaign.ID, pageParams); page.Size != 1 {
		t.Fatalf("got unexpected number of referrals for campaign")
	}
	if page := referralRepo.GetReferralsWithStatus(domain.PendingStatus, pageParams); page.Size != 1 {
		t.Fatalf("got unexpected number of pending referrals")
	}

	// Update (set status)
	updated, err := referralRepo.UpdateReferral(referral.ID, domain.VerifiedStatus)
	if err != nil {
		t.Fatalf("failed to update referral status: %+v", err)
	}
	if updated.Status != domain.VerifiedStatus {
		t.Fatalf("expected verified status, got: %s", updated.Status)
	}
	if page := referralRepo.GetReferralsWithStatus(domain.VerifiedStatus, pageParams); page.Size != 1 {
		t.Fatalf("got unexpected number of verified referrals")
	}

	// Error check: ensure accounts can't add referrals for thier own campaigns.
	if _, err := referralRepo.CreateReferral(campaign.ID, referer); err == nil {
		t.Fatalf("expected self referral error")
	}
}
