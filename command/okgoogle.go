package command

type OKGoogle struct {
	command Command
}

func NewOKGoogle() *OKGoogle {
	return &OKGoogle{}
}

func (o *OKGoogle) SetCommand(command Command) {
	o.command = command
}

func (o *OKGoogle) Talk() {
	o.command.Run()
}
