package stopwatch

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/zlypher/go-timer/command"
	"github.com/zlypher/go-timer/timeutil"
)

var in io.Reader = os.Stdin
var out io.Writer = os.Stdout
var isRunning bool = true

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
		return command.RESULT_INTERNAL_ERROR
	}

	// Check for input in the background
	go func() {
		scanner := bufio.NewScanner(in)

		for scanner.Scan() {
			text := scanner.Text()
			if text == "q" {
				isRunning = false
				break
			}

			fmt.Fprintf(out, "INPUT: %v\n", text)
		}
	}()

	runStopwatch(interval)
	return command.RESULT_SUCCESS
}

// runStopwatch runs the stopwatch and prints the elapsed time.
func runStopwatch(interval time.Duration) {
	startTime := time.Now()

	for _ = range time.Tick(interval) {
		if !isRunning {
			break
		}

		elapsed := time.Since(startTime)
		fmt.Fprintf(out, "\r%v", timeutil.Format(elapsed))
	}
}
