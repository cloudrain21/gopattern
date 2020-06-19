package command

import "fmt"

type StringCommand struct {
	v string
}

func (s StringCommand)Execute() {
	fmt.Println(s.v)
}

func NewStringCommand(a string) Command {
	return &StringCommand{a}
}

