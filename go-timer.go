package main

import (
	"fmt"
	"io"
	"os"

	"github.com/zlypher/go-timer/command"
	"github.com/zlypher/go-timer/countdown"
	"github.com/zlypher/go-timer/interval"
	"github.com/zlypher/go-timer/stopwatch"
)

var out io.Writer = os.Stdout

func main() {
	result := runGoTimer(os.Args)
	os.Exit(result)
}

// runGoTimer runs the go-timer app with the given arguments.
func runGoTimer(args []string) int {
	// Sanity check arguments
	if len(args) < 2 {
		printUsage()
		return command.RESULT_MISSING_ARGS
	}

	name := args[1]
	arguments := args[2:]

	commands := setupCommands()
	cmd, ok := commands[name]
	if !ok {
		fmt.Fprintf(out, "Couldn't find command: %v\n", name)
		printUsage()
		return command.RESULT_CMD_NOT_FOUND
	}

	return cmd.Run(arguments)
}

// setupCommands prepares a mapping of the available commands.
func setupCommands() command.CommandMap {
	commands := make(command.CommandMap)
	commands["countdown"] = countdown.Countdown{}
	commands["stopwatch"] = stopwatch.Stopwatch{}
	commands["interval"] = interval.Interval{}
	commands["version"] = command.CallCommand{Func: printVersion}
	commands["usage"] = command.CallCommand{Func: printUsage}
	commands["help"] = command.CallCommand{Func: printUsage}

	return commands
}

// printUsage prints a help message for go-timer
func printUsage() {
	fmt.Fprintln(out, "go-timer - Timer cool written in GO")
	fmt.Fprintf(out, "------------------------------------\n\n")
	fmt.Fprintf(out, "Usage: go-timer <command> \n\n")
	fmt.Fprintf(out, "where <command> is one of:\n")
	fmt.Fprintf(out, "\tcountdown, interval, stopwatch, version, help, usage\n")
}

// printVersion prints the current go-timer version
func printVersion() {
	fmt.Fprintln(out, "go-timer 0.0.1")
}
