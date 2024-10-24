package handler

import (
	"log"
	"os"
)

// Defaults for redirect config.
const defaultCookieMaxAge = 30 * 24 * 60 * 60
const defaultCookieName = "_referral_campaign"

// RedirectConfig holds URLs for the redirect handler.
type RedirectConfig struct {
	SignupURL    string
	TargetURL    string
	CookieMaxAge int
	CookieName   string
}

// LoadRedirectConfig loads config for a redirect handler from env vars.
func LoadRedirectConfig() RedirectConfig {
	return RedirectConfig{
		SignupURL:    requireEnvVar("SIGNUP_URL"),
		TargetURL:    requireEnvVar("TARGET_URL"),
		CookieMaxAge: defaultCookieMaxAge,
		CookieName:   defaultCookieName,
	}
}

// Lookup target URL from env var and panic if not found.
func requireEnvVar(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("%n not defined", key)
	}
	return value
}
