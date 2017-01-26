package countdown

import (
	"fmt"
	"os"
	"time"
)

// Run sets up the countdown and starts it.
func Run(args []string) {
	// Initialize countdown
	input := args[0]
	duration, err := time.ParseDuration(input)
	if err != nil {
		fmt.Printf("Failed to parse input: %v", err)
		os.Exit(1)
	}
	interval, err := time.ParseDuration("10ms")
	if err != nil {
		fmt.Printf("Failed to parse interval: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Counting down %v\n", duration)

	runCountdown(duration, interval)
}

// runCountdown counts down from the given duration in the given interval steps.
func runCountdown(duration, interval time.Duration) {
	endTime := time.Now().Add(duration)

	for tick := range time.Tick(interval) {
		// If countdown reached the end
		if tick.After(endTime) {
			fmt.Printf("\r00:00:00.000")
			fmt.Printf("\a")
			return
		}

		diff := endTime.Sub(tick)
		printRemainingTime(diff)
	}
}

// printRemainingTime prints the remaining time based on the given duration.
func printRemainingTime(dur time.Duration) {
	ns := dur.Nanoseconds()
	timestamp := time.Unix(0, ns).UTC()
	fmt.Printf("\r%v", timestamp.Format("15:04:05.000"))
}
