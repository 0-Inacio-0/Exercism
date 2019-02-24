// This package calculate the moment when someone has lived for 10^9 seconds.
package gigasecond

import "time"

const (
	// Gigasecond is 10 ^ 9 seconds
	Gigasecond = 1e9 * time.Second
)

//This function receives a date and time and returns it added a gigasecond(10^9 seconds)
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
