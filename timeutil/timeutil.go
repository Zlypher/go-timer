package timeutil

import (
	"fmt"
	"time"
)

// Print prints the given duration in the format "hh:mm:ss.ms".
func Print(duration time.Duration) string {
	ns := duration.Nanoseconds()
	timestamp := time.Unix(0, ns).UTC()

	return timestamp.Format("15:04:05.000")
}

// printAlternative is an alternative implementation of the Print method.
func printAlternative(duration time.Duration) string {
	ns := duration.Nanoseconds()
	hh := ns / (1000 * 1000 * 1000 * 60 * 60)
	mm := ns / (1000 * 1000 * 1000 * 60) % 60
	ss := ns / (1000 * 1000 * 1000) % 60
	ms := ns / (1000 * 1000) % 1000

	return fmt.Sprintf("%02d:%02d:%02d.%03d", hh, mm, ss, ms)
}
