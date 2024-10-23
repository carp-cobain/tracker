package handler

import (
	"log"
	"os"
)

type RedirectConfig struct {
	SignupURL string
	TargetURL string
}

func DefaultRedirectConfig() RedirectConfig {
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
