package command

import "testing"

func TestCallCommand(t *testing.T) {
	called := false
	fn := func() { called = true }

	cmd := CallCommand{Func: fn}
	cmd.Run([]string{})

	if !called {
		t.Errorf("CallCommand.Run(fn) didn't call fn")
	}
}
