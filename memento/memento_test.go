package memento

func ExampleMemento() {
	savedStatus := make([]*Memento,0)

	originator := NewOriginator()
	originator.Set("state1")
	originator.Set("state2")
	savedStatus = append(savedStatus, originator.SaveToMemento())
	originator.Set("state3")
	savedStatus = append(savedStatus, originator.SaveToMemento())
	originator.Set("state4")
	originator.RestoreFromMemento(savedStatus[1])

	// Output:
	// Originator: Setting state to  state1
	// Originator: Setting state to  state2
	// Originator: Saving to Memento.
	// Originator: Setting state to  state3
	// Originator: Saving to Memento.
	// Originator: Setting state to  state4
	// Originator : State after restoring from memento :  state3
}
