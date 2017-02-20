package command

// Commander defines an interface for a command to run.
type Commander interface {
	Description() string
	Run(args []string) int
}

type CommandMap map[string]Commander

type CallFn func()

// CallCommand is a command to simply call a given function.
type CallCommand struct {
	Func CallFn
}

func (c CallCommand) Description() string {
	return "Not Set"
}

// Run runs the CallCommand to simply call the given function.
func (c CallCommand) Run(args []string) int {
	c.Func()
	return 0
}
