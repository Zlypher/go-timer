# go-timer

[![CircleCI](https://circleci.com/gh/zlypher/go-timer.svg?style=shield)](https://circleci.com/gh/zlypher/go-timer)

go-timer is a simple cli countdown timer. It is written in [Golang](https://golang.org/) (mainly to get a better understanding of the language). Feel free to fork/modify/play around with it as you see fit.

It currently uses the string representation of time.Duration to specfiy the time to count down. So for example "5m30s" would result in a countdown timer of 5 minutes and 30 seconds.

# Usage

```
go-timer [command] [options]
```

## Commands

### stopwatch

Starts a stopwatch timer. Currently the stopwatch counts until the command is cancelled.

```bash
go-timer stopwatch # Starts the stopwatch
```

### countdown

Starts a countdown timer with the specified duration. Currently only values are supported, that can be natively parsed by [golang's time.ParseDuration](https://golang.org/pkg/time/#ParseDuration)

```bash
go-timer countdown 10s # Countdown from 10 seconds
go-timer countdown 5m # Countdown from 5 minutes
```

### interval

Starts an interval timer with the specified durations.

```bash
go-timer interval 5s 2m 3s # counts from 0 to 5 seconds, then from 0 to 2 minutes, then from 0 to 3 seconds
```

### version

Prints the current version of go-timer.

```
go-timer version
```

### help/usage

Prints a help message for the general usage of go-timer.

```
go-timer help
go-timer usage
```

# LICENSE

[MIT](https://choosealicense.com/licenses/mit/)
