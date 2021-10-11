package command

import "fmt"

type IntCommand struct {
	v int
}

func (i IntCommand) Execute() {
	fmt.Println(i.v)
}

func NewIntCommand(a int) Command {
	return &IntCommand{a}
}
