package command

func ExampleCommand() {
	commandList := []Command{}

	commandList = append(commandList, NewIntCommand())
	commandList = append(commandList, NewStringCommand())
	commandList = append(commandList, NewIntCommand())

	for _, c := range commandList {
		c.Execute()
	}

	// Output:
	// Int Command
	// String Command
	// Int Command
}
