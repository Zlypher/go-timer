package main

import (
	"fmt"
	"os"

	"github.com/zlypher/go-timer/countdown"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "countdown":
		countdown.Run(os.Args[2:])
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

// printUsage prints a help message for go-timer
func printUsage() {
	fmt.Println("go-timer - Timer cool written in GO")
	fmt.Printf("Usage: %s countdown \n\n", os.Args[0])
}

// printVersion prints the current go-timer version
func printVersion() {
	fmt.Println("go-timer 0.0.1")
}
