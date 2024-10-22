package dto

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// ensure prefix is loaded only once
var once sync.Once

// account address prefix to be set from env var
var accountPrefix string

// validate a blockchain account address
func ValidateAccount(account string) (string, error) {
	// Ensure prefix is loaded from env only once
	once.Do(loadPrefixEnv)
	account = strings.TrimSpace(account)
	if account == "" {
		return "", fmt.Errorf("account address cannot be blank")
	}
	if strings.ToLower(account) != account {
		return "", fmt.Errorf("account address must be lower case")
	}
	if !strings.HasPrefix(account, accountPrefix) {
		return "", fmt.Errorf("account address must have prefix: %s", accountPrefix)
	}
	length := len(account)
	if length < 41 || length > 61 {
		return "", fmt.Errorf("invalid account address length: %d", length)
	}
	return account, nil
}

// Load account address prefix based on env var
func loadPrefixEnv() {
	if mainnet := os.Getenv("MAINNET"); mainnet == "true" {
		accountPrefix = "pb"
	} else {
		accountPrefix = "tp"
	}
}
