package countdown

import (
	"fmt"
	"os"
	"time"

	"github.com/zlypher/go-timer/timeutil"
)

type Countdown struct{}

func (c Countdown) Description() string {
	return "Hello countdown"
}

// Run sets up the countdown and starts it.
func (c Countdown) Run(args []string) {
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
		fmt.Printf("\r%v", timeutil.Format(diff))
	}
}
