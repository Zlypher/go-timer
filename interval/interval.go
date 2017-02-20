package interval

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/zlypher/go-timer/timeutil"
)

var out io.Writer = os.Stdout

type intervalConfig struct {
	intervals []time.Duration
	repeat    int
}

type Interval struct{}

func (i Interval) Description() string {
	return "Interval impl"
}

func (i Interval) Run(args []string) int {
	// TODO: Move the flag set somewhere else at a alter point
	intervalCommand := flag.NewFlagSet("interval", flag.ExitOnError)
	repeatPtr := intervalCommand.Int("repeat", 1, "Number the intervals should be repeated.")
	intervalCommand.Parse(args)

	fmt.Fprintln(out, "Starting interval")

	interval, err := time.ParseDuration("10ms")
	if err != nil {
		fmt.Fprintf(out, "Failed to parse interval: %v", err)
		return 1
	}

	// Parse arguments into config as there might be invalid values
	config := parseArguments(args)
	config.repeat = *repeatPtr

	runIntervals(config, interval)
	return 0
}

// parseArguments parses the given cmd arguments into an intervalConfig.
func parseArguments(args []string) intervalConfig {
	config := intervalConfig{}

	for _, arg := range args {
		parsed, err := time.ParseDuration(arg)
		if err == nil {
			config.intervals = append(config.intervals, parsed)
		}
	}

	return config
}

// runIntervals runs the various intervals defined in the given configuration.
func runIntervals(config intervalConfig, interval time.Duration) {
	for i := 0; i < config.repeat; i++ {
		for _, in := range config.intervals {
			start := time.Now()
			ticker := time.NewTicker(interval)

			for _ = range ticker.C {
				elapsed := time.Since(start)
				fmt.Fprintf(out, "\r%v", timeutil.Format(elapsed))

				if elapsed > in {
					ticker.Stop()
					fmt.Fprintln(out)
					fmt.Fprintln(out, "-----")
					break
				}
			}
		}
	}
}
