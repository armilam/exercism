// An exercism exercise to learn time
package gigasecond

// import path for the time package from the standard library
import "time"

// Adds a billion seconds to the given time
func AddGigasecond(t time.Time) time.Time {
	gigasecond, _ := time.ParseDuration("1000000000s")
	return t.Add(gigasecond)
}
