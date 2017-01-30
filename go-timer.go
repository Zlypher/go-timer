package main

import (
	"fmt"
	"os"

	"github.com/zlypher/go-timer/command"
	"github.com/zlypher/go-timer/countdown"
	"github.com/zlypher/go-timer/stopwatch"
)

func main() {
	// Sanity check arguments
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	commands := setupCommands()

	// Determine command
	switch os.Args[1] {
	case "countdown":
		commands["countdown"].Run(os.Args[2:])
	case "stopwatch":
		commands["stopwatch"].Run(os.Args[2:])
	case "version":
		printVersion()
		os.Exit(0)
	case "usage":
		fallthrough
	case "help":
		fallthrough
	default:
		printUsage()
		os.Exit(1)
	}
}

// setupCommands prepares a mapping of the available commands.
func setupCommands() command.CommandMap {
	commands := make(command.CommandMap)
	commands["countdown"] = countdown.Countdown{}
	commands["stopwatch"] = stopwatch.Stopwatch{}
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
