package main

import (
	"fmt"
	"os"

	"github.com/zlypher/go-timer/command"
	"github.com/zlypher/go-timer/countdown"
	"github.com/zlypher/go-timer/interval"
	"github.com/zlypher/go-timer/stopwatch"
)

func main() {
	args := os.Args
	// Sanity check arguments
	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	name := args[1]
	arguments := args[2:]

	commands := setupCommands()
	command, ok := commands[name]
	if !ok {
		fmt.Printf("Couldn't find command: %v\n", name)
		printUsage()
		os.Exit(1)
	}

	command.Run(arguments)
	os.Exit(0)
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
	fmt.Println("go-timer - Timer cool written in GO")
	fmt.Printf("Usage: %s countdown \n\n", os.Args[0])
}

// printVersion prints the current go-timer version
func printVersion() {
	fmt.Println("go-timer 0.0.1")
}
