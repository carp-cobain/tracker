package dto

import (
	"fmt"
	"os"
	"strings"
)

// validate a blockchain account address
func ValidateAccount(account string) (string, error) {
	account = strings.TrimSpace(account)
	if account == "" {
		return "", fmt.Errorf("account address cannot be blank")
	}
	if strings.ToLower(account) != account {
		return "", fmt.Errorf("account address must be lower case")
	}
	prefix := lookupPrefixEnv()
	if !strings.HasPrefix(account, prefix) {
		return "", fmt.Errorf("account address must have prefix: %s", prefix)
	}
	length := len(account)
	if length < 41 || length > 61 {
		return "", fmt.Errorf("invalid account address length: %d", length)
	}
	return account, nil
}

// Determine address prefix based on env var
func lookupPrefixEnv() string {
	if mainnet, ok := os.LookupEnv("MAINNET"); ok {
		if mainnet == "true" {
			return "pb"
		}
	}
	return "tp"
}
