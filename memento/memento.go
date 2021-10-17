package memento

type Memento struct {
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{
		state: state,
	}
}

func (m *Memento) GetState() string {
	return m.state
}
