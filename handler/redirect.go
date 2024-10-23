package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/carp-cobain/tracker/domain"
	"github.com/carp-cobain/tracker/dto"
	"github.com/carp-cobain/tracker/keeper"

	"github.com/gin-gonic/gin"
)

// CookieName is the name for referral campaign cookies
var CookieName string = "_referral_campaign"

// MaxAge is the max age for referral campaign cookies
var MaxAge int = 30 * 24 * 60 * 60

// RedirectHandler is the http/json api for managing referral campaigns
type RedirectHandler struct {
	campaignReader keeper.CampaignReader
	referralKeeper keeper.ReferralKeeper
}

// NewRedirectHandler creates a new campaign redirect handler
func NewRedirectHandler(
	campaignReader keeper.CampaignReader,
	referralKeeper keeper.ReferralKeeper,
) RedirectHandler {
	return RedirectHandler{campaignReader, referralKeeper}
}

// GET /tracker/referrals/signup
// SignupRedirect drops a referral campaign cookie and redirects to a signup URL.
func (self RedirectHandler) SignupRedirect(c *gin.Context) {
	signupURL := envSignupURL()
	campaign := c.Query("campaign")
	if campaign == "" {
		log.Println("no campaign param found")
		c.Redirect(http.StatusFound, signupURL)
		return
	}
	campaignID, err := strconv.ParseUint(campaign, 10, 64)
	if err != nil {
		log.Printf("failed to parse campaign param: %s", err.Error())
		c.Redirect(http.StatusFound, signupURL)
		return
	}
	if campaign, err := self.campaignReader.GetCampaign(campaignID); err == nil {
		if campaign.Type == "referral" {
			c.SetCookie(
				CookieName,
				fmt.Sprintf("%d", campaign.ID),
				min(MaxAge, campaign.TTL()),
				os.Getenv("SIGNUP_COOKIE_PATH"),
				os.Getenv("SIGNUP_COOKIE_DOMAIN"),
				false,
				false,
			)
		}
	}
	c.Redirect(http.StatusFound, signupURL)
}

// GET /tracker/referrals
// ReferralCaptureRedirect records signup referrals from campaign cookies then redirects to a target URL.
func (self RedirectHandler) ReferralCaptureRedirect(c *gin.Context) {
	// Use env var target URL for redirect
	targetURL := envTargetURL()
	// Lookup campaign from cookie
	campaign, err := self.cookieCampaign(c)
	if err != nil {
		log.Printf("failed to lookup campaign from cookie: %s", err.Error())
		c.Redirect(http.StatusFound, targetURL)
		return
	}
	// Assumes blockchain address is created during signup
	account, err := self.findAccount(c)
	if err != nil {
		log.Printf("no valid blockchain account address found: %s", err.Error())
		c.Redirect(http.StatusFound, targetURL)
		return
	}
	// Store referral
	if campaign.Type == "referral" {
		if _, err := self.referralKeeper.CreateReferral(campaign.ID, account); err != nil {
			log.Printf("failed to record referral: %s", err.Error())
		}
	}
	// Send user on their way
	c.Redirect(http.StatusFound, targetURL)
}

// get referral campaign using cookie set during signup redirect.
func (self RedirectHandler) cookieCampaign(c *gin.Context) (campaign domain.Campaign, err error) {
	// Check for cookie, redirect if not found.
	var cookie string
	if cookie, err = c.Cookie(CookieName); err != nil {
		return
	}
	// Lookup campaign for cookie value
	var id uint64
	if id, err = strconv.ParseUint(cookie, 10, 64); err == nil {
		campaign, err = self.campaignReader.GetCampaign(id)
	}
	return
}

// find blockchain account address from header or query param
func (self RedirectHandler) findAccount(c *gin.Context) (string, error) {
	account := c.GetHeader("x-account-address")
	if account == "" {
		account = c.Query("account")
	}
	account, err := dto.ValidateAccount(account)
	if err != nil {
		return "", err
	}
	return account, nil
}

// Lookup signup URL from env var
func envSignupURL() string {
	signupURL, ok := os.LookupEnv("SIGNUP_URL")
	if !ok {
		log.Panicf("SIGNUP_URL not defined")
	}
	return signupURL
}

// Lookup target URL from env var
func envTargetURL() string {
	targetURL, ok := os.LookupEnv("TARGET_URL")
	if !ok {
		log.Panicf("TARGET_URL not defined")
	}
	return targetURL
}
