package command

type HeaterOnCommand struct {
	heater *Heater
}

func NewHeaterOnCommand(heater *Heater) *HeaterOnCommand {
	return &HeaterOnCommand{
		heater: heater,
	}
}

func (h *HeaterOnCommand) Run() {
	h.heater.powerOn()
}
