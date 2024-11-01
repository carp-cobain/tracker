package domain

import (
	"fmt"
	"os"
	"strings"
)

// Account is a blockchain account address
type Account struct {
	address string
}

// NewAccount creates a new account with the given address.
func NewAccount(address string) Account {
	return Account{address}
}

// MustValidateAccount creates a new account and panics if the account address is invalid.
func MustValidateAccount(address string) Account {
	return Must(NewAccount(address).Validate())
}

// Must returns the account or panics if the given err is not nil.
func Must(account Account, err error) Account {
	if err != nil {
		panic(err)
	}
	return account
}

// Validate checks whether an account's blockchain address is valid.
func (self Account) Validate() (Account, error) {
	address := strings.TrimSpace(self.address)
	if address == "" {
		return self, fmt.Errorf("account address cannot be blank")
	}
	if strings.ToLower(address) != address {
		return self, fmt.Errorf("account address must be lower case")
	}
	addressPrefix := loadPrefixEnv()
	if !strings.HasPrefix(address, addressPrefix) {
		return self, fmt.Errorf("account address must have prefix: %s", addressPrefix)
	}
	length := len(address)
	if length < 41 || length > 61 {
		return self, fmt.Errorf("invalid account address length: %d", length)
	}
	return self, nil
}

// String returns the account address.
func (self Account) String() string {
	return self.address
}

// MarshalText implements encoding.TextMarshaler.
func (self Account) MarshalText() ([]byte, error) {
	return []byte(self.address), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (self *Account) UnmarshalText(data []byte) (err error) {
	*self, err = NewAccount(string(data)).Validate()
	return
}

// Load account address prefix based on env var
func loadPrefixEnv() string {
	if mainnet := os.Getenv("MAINNET"); mainnet == "true" {
		return "pb"
	}
	return "tp"
}
