// Package iso6801 contains a custom time type which
// (de)serializes in ISO8601 format.
//
// see: https://en.wikipedia.org/wiki/ISO_8601.
package iso8601

import (
	"strings"
	"time"
)

const Layout = "2006-01-02T15:04:05.999Z"

type Time struct {
	time.Time
}

// Parse a string into an ISO8601 time.
// Returns an error if the string is not in ISO8601 format.
func Parse(value string) (Time, error) {
	nt, err := time.Parse(Layout, value)
	if err != nil {
		return Time{}, err
	}
	return Time{nt}, nil
}

// UnmarshalJSON Parses the json string in the custom format
func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)

	nt, err := time.Parse(Layout, s)

	*ct = Time{nt}
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct Time) MarshalJSON() ([]byte, error) {
	str := `"` + ct.String() + `"`
	return []byte(str), nil
}

// String returns the time in ISO8601 format
func (ct Time) String() string {
	return ct.Time.UTC().Format(Layout)
}

// New creates a new ISO8601 Time from a standard Go time.
func New(t time.Time) Time {
	return Time{t}
}

// Now is a helper to create a ISO8601 Time at the current instant.
func Now() Time {
	return Time{time.Now()}
}
