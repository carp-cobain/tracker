package handler

import (
	"log"
	"os"
)

// RedirectConfig holds URLs for the redirect handler.
type RedirectConfig struct {
	SignupURL string
	TargetURL string
}

// LoadRedirectConfig loads config for a redirect handler from env vars.
func LoadRedirectConfig() RedirectConfig {
	return RedirectConfig{
		SignupURL: requireEnvVar("SIGNUP_URL"),
		TargetURL: requireEnvVar("TARGET_URL"),
	}
}

// Lookup target URL from env var and panic if not found.
func requireEnvVar(name string) string {
	url, ok := os.LookupEnv(name)
	if !ok {
		log.Panicf("%n not defined", name)
	}
	return url
}
