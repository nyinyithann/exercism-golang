// Package gigasecond contains a function to calculate gigasecond time for a person
package gigasecond

import "time"

// AddGigasecond calculates gigasecond time for a person
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
