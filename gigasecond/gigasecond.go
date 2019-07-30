// Package gigasecond calculates the moment someone have lived for 10^9 seconds
package gigasecond

import "time"

const gigasecond time.Duration = 1e9 * time.Second

// AddGigasecond returns t after a gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
