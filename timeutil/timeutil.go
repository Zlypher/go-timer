package timeutil

import (
	"fmt"
	"time"
)

// Format formats the given duration in the format "hh:mm:ss.ms".
func Format(duration time.Duration) string {
	ns := duration.Nanoseconds()
	timestamp := time.Unix(0, ns).UTC()

	return timestamp.Format("15:04:05.000")
}

// formatAlternative is an alternative implementation of the Format method.
func formatAlternative(duration time.Duration) string {
	ns := duration.Nanoseconds()
	hh := ns / (1000 * 1000 * 1000 * 60 * 60)
	mm := ns / (1000 * 1000 * 1000 * 60) % 60
	ss := ns / (1000 * 1000 * 1000) % 60
	ms := ns / (1000 * 1000) % 1000

	return fmt.Sprintf("%02d:%02d:%02d.%03d", hh, mm, ss, ms)
}
