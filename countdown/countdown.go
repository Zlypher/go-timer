package countdown

import (
	"fmt"
	"time"
)

func Run(args []string) error {
	// Initialize countdown
	input := args[0]
	duration, err := time.ParseDuration(input)
	if err != nil {
		return err
	}
	interval, err := time.ParseDuration("10ms")
	if err != nil {
		return err
	}

	fmt.Printf("Counting down %v\n", duration)

	runCountdown(duration, interval)
	return nil
}

func runCountdown(duration, interval time.Duration) {
	endTime := time.Now().Add(duration)

	for tick := range time.Tick(interval) {
		diff := endTime.Sub(tick)
		if tick.After(endTime) {
			fmt.Printf("\r00:00:00.000")
			return
		}

		print(diff)
	}
}

func print(dur time.Duration) {
	ns := dur.Nanoseconds()
	timestamp := time.Unix(0, ns).UTC()
	fmt.Printf("\r%v", timestamp.Format("15:04:05.000"))
}
