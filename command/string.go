package command

import "fmt"

type StringCommand struct{}

func (i *StringCommand) Execute() {
	fmt.Println("String Command")
}

func NewStringCommand() Command {
	return &StringCommand{}
}
