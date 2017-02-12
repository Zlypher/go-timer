package stopwatch

import (
	"fmt"
	"os"
	"time"

	"github.com/zlypher/go-timer/timeutil"
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

	// Setup something to read and react to input

	runStopwatch(interval)
}

// runStopwatch runs the stopwatch and prints the elapsed time.
func runStopwatch(interval time.Duration) {
	startTime := time.Now()

	for _ = range time.Tick(interval) {
		elapsed := time.Since(startTime)
		fmt.Printf("\r%v", timeutil.Format(elapsed))
	}
}
