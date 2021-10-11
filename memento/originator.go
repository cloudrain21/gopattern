package memento

import "fmt"

type Originator struct {
	state string
}

func NewOriginator() *Originator {
	return &Originator{}
}

func (o *Originator)Set(state string) {
	o.state = state
	fmt.Println("Originator: Setting state to ", state)
}

func (o *Originator)SaveToMemento() *Memento{
	fmt.Println("Originator: Saving to Memento.")
	return NewMemento(o.state)
}

func (o *Originator)RestoreFromMemento(m *Memento) {
	o.state = m.GetSavedState()
	fmt.Println("Originator : State after restoring from memento : ", o.state)
}