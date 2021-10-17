package command

type LampOnCommand struct {
	lamp *Lamp
}

func NewLampOnCommand(lamp *Lamp) *LampOnCommand {
	return &LampOnCommand{
		lamp: lamp,
	}
}

func (l *LampOnCommand) Run() {
	l.lamp.turnOn()
}
