package interval

import (
	"fmt"
	"os"
	"time"

	"github.com/zlypher/go-timer/timeutil"
)

type intervalConfig struct {
	intervals []time.Duration
	repeat    bool
}

type Interval struct{}

func (i Interval) Description() string {
	return "Interval impl"
}

func (i Interval) Run(args []string) {
	fmt.Println("Starting interval")

	interval, err := time.ParseDuration("10ms")
	if err != nil {
		fmt.Printf("Failed to parse interval: %v", err)
		os.Exit(1)
	}

	// Parse arguments into config as there might be invalid values
	config := parseArguments(args)

	runIntervals(config, interval)
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
	for _, in := range config.intervals {
		start := time.Now()
		ticker := time.NewTicker(interval)

		for _ = range ticker.C {
			elapsed := time.Since(start)
			fmt.Printf("\r%v", timeutil.Format(elapsed))

			if elapsed > in {
				ticker.Stop()
				fmt.Println()
				fmt.Println("-----")
				break
			}
		}
	}
}
