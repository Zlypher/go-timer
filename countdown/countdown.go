package countdown

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/zlypher/go-timer/command"
	"github.com/zlypher/go-timer/timeutil"
)

var out io.Writer = os.Stdout

type Countdown struct{}

func (c Countdown) Description() string {
	return "Hello countdown"
}

// Run sets up the countdown and starts it.
func (c Countdown) Run(args []string) int {
	// Initialize countdown
	input := args[0]
	duration, err := time.ParseDuration(input)
	if err != nil {
		fmt.Fprintf(out, "Failed to parse input: %v", err)
		return command.RESULT_INVALID_ARG
	}
	interval, err := time.ParseDuration("10ms")
	if err != nil {
		fmt.Fprintf(out, "Failed to parse interval: %v", err)
		return command.RESULT_INTERNAL_ERROR
	}

	fmt.Fprintf(out, "Counting down %v\n", duration)

	runCountdown(duration, interval)
	return command.RESULT_SUCCESS
}

// runCountdown counts down from the given duration in the given interval steps.
func runCountdown(duration, interval time.Duration) {
	endTime := time.Now().Add(duration)

	for tick := range time.Tick(interval) {
		// If countdown reached the end
		if tick.After(endTime) {
			fmt.Fprintf(out, "\r00:00:00.000")
			fmt.Fprintf(out, "\a")
			return
		}

		diff := endTime.Sub(tick)
		fmt.Fprintf(out, "\r%v", timeutil.Format(diff))
	}
}
