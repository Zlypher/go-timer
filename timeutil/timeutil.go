package timeutil

import "time"

// Print prints the given duration in the format "hh:mm:ss.ms".
func Print(duration time.Duration) string {
	ns := duration.Nanoseconds()
	timestamp := time.Unix(0, ns).UTC()

	return timestamp.Format("15:04:05.000")
}
