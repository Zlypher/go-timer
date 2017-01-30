package stopwatch

import (
	"fmt"
	"os"
	"time"
)

type Stopwatch struct{}

func (s Stopwatch) Description() string {
	return "Stopwatch impl"
}

// Run starts the stopwatch.
func (s Stopwatch) Run(args []string) {
	fmt.Println("Starting stopwatch")

	interval, err := time.ParseDuration("10ms")
	if err != nil {
		fmt.Printf("Failed to parse interval: %v", err)
		os.Exit(1)
	}

	runStopwatch(interval)
}

// runStopwatch runs the stopwatch and prints the elapsed time.
func runStopwatch(interval time.Duration) {
	startTime := time.Now()

	for _ = range time.Tick(interval) {
		elapsed := time.Since(startTime)
		printTime(elapsed)
	}
}

// printTime prints the remaining time based on the given duration.
func printTime(dur time.Duration) {
	ns := dur.Nanoseconds()
	timestamp := time.Unix(0, ns).UTC()
	fmt.Printf("\r%v", timestamp.Format("15:04:05.000"))
}
