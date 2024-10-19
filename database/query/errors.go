package query

import "errors"

// ErrCampaignExpired is an error indicating a campaign has expired.
var ErrCampaignExpired = errors.New("campaign has expired")
