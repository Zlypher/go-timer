package command

type Commander interface {
	Description() string
	Run(args []string)
}
