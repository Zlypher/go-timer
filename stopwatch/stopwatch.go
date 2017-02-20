package stopwatch

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/zlypher/go-timer/timeutil"
)

var out io.Writer = os.Stdout

type Stopwatch struct{}

func (s Stopwatch) Description() string {
	return "Stopwatch impl"
}

// Run starts the stopwatch.
func (s Stopwatch) Run(args []string) int {
	fmt.Fprintln(out, "Starting stopwatch")

	interval, err := time.ParseDuration("10ms")
	if err != nil {
		fmt.Fprintf(out, "Failed to parse interval: %v", err)
		return 1
	}

	// Setup something to read and react to input

	runStopwatch(interval)
	return 0
}

// runStopwatch runs the stopwatch and prints the elapsed time.
func runStopwatch(interval time.Duration) {
	startTime := time.Now()

	for _ = range time.Tick(interval) {
		elapsed := time.Since(startTime)
		fmt.Fprintf(out, "\r%v", timeutil.Format(elapsed))
	}
}
