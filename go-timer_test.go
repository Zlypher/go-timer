package main

import (
	"bytes"
	"testing"

	"github.com/zlypher/go-timer/command"
)

func TestSetupCommands(t *testing.T) {
	requiredCommands := []string{"stopwatch", "countdown", "interval", "version", "usage", "help"}
	requiredLength := len(requiredCommands)
	commands := setupCommands()

	if len(commands) != requiredLength {
		t.Errorf("len(%v) == %d, want %d", commands, len(commands), requiredLength)
	}

	for _, k := range requiredCommands {
		cmd, ok := commands[k]
		if !ok {
			t.Errorf("commands[%v] == %v, expected to find command", k, cmd)
		}
	}
}

func TestRunGoTimer(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	cases := []struct {
		in   []string
		want int
	}{
		{nil, command.RESULT_MISSING_ARGS},
		{[]string{}, command.RESULT_MISSING_ARGS},
		{[]string{"go-timer"}, command.RESULT_MISSING_ARGS},
		{[]string{"go-timer", "version"}, command.RESULT_SUCCESS},
		{[]string{"go-timer", "INVALID"}, command.RESULT_CMD_NOT_FOUND},
	}

	for _, c := range cases {
		result := runGoTimer(c.in)

		if result != c.want {
			t.Errorf("runGoTimer(%v) == %d, want %d", c.in, result, c.want)
		}
	}
}

func TestPrintVersion(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	expected := "go-timer 0.0.1\n"
	printVersion()
	actual := out.(*bytes.Buffer).String()
	if actual != expected {
		t.Errorf("printVersion() prints '%v', expected '%v'", actual, expected)
	}
}
