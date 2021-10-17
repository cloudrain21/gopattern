package command

func ExampleCommand() {
	heater := NewHeater()
	lamp := NewLamp()

	heaterOnCommand := NewHeaterOnCommand(heater)
	lampOnCommand := NewLampOnCommand(lamp)

	okGoogle := NewOKGoogle()

	okGoogle.SetCommand(heaterOnCommand)
	okGoogle.Talk()

	okGoogle.SetCommand(lampOnCommand)
	okGoogle.Talk()

	// Output:
	// Heater on
	// Lamp on
}
