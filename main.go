package main

import (
	"log"
	"os"

	"github.com/carp-cobain/tracker/database"
	"github.com/carp-cobain/tracker/database/repo"
	"github.com/carp-cobain/tracker/handler"
	"github.com/carp-cobain/tracker/processor"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	// Perform env checks
	if _, ok := os.LookupEnv("DISABLE_COLOR"); ok {
		gin.DisableConsoleColor()
	}

	// DB
	readDB, writeDB, err := database.ConnectAndMigrate()
	if err != nil {
		log.Panicf("unable to connnect to db: %+v", err)
	}

	// Repos
	campaignRepo := repo.NewCampaignRepo(readDB, writeDB)
	referralRepo := repo.NewReferralRepo(readDB, writeDB)

	// Handlers
	campaignHandler := handler.NewCampaignHandler(campaignRepo)
	referralHandler := handler.NewReferralHandler(campaignRepo, referralRepo)

	// API
	r := gin.Default()
	v1 := r.Group("/tracker/api/v1")
	{
		v1.GET("/campaigns", campaignHandler.GetCampaigns)
		v1.POST("/campaigns", campaignHandler.CreateCampaign)
		v1.GET("/campaigns/:id", campaignHandler.GetCampaign)
		v1.DELETE("/campaigns/:id", campaignHandler.ExpireCampaign)
		v1.GET("/campaigns/:id/referrals", referralHandler.GetReferrals)
		v1.POST("/campaigns/:id/referrals", referralHandler.CreateReferral)
	}

	// Run background processors
	referralVerifier := processor.NewReferralVerifier(referralRepo)
	c := cron.New()
	c.AddFunc("@every 1m", referralVerifier.VerifyReferrals)
	c.Start()

	// Run server
	if err := r.Run(); err != nil {
		log.Panicf("unable to start tracker server:  %+v", err)
	}
}
