package model

import "time"

// DateTime is used to store timestamps as INT columns in SQLite
type DateTime int64

// ToDomain creates a go stdlib time from the SQLite column type (unix seconds).
func (t *DateTime) ToDomain() time.Time {
	return time.Unix(int64(*t), 0).UTC()
}

// Now returns the current time as a SQLite column type.
func Now() DateTime {
	return DateTime(time.Now().Unix())
}

// Expiry returns one year from now as a SQLite column type.
func Expiry() DateTime {
	return DateTime(time.Now().Add(365 * 24 * time.Hour).Unix())
}
