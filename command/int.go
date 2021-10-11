package command

import "fmt"

type IntCommand struct{}

func (i *IntCommand) Execute() {
	fmt.Println("Int Command")
}

func NewIntCommand() Command {
	return &IntCommand{}
}
