package memento

func ExampleMemento() {
	savedStatus := make([]*Memento, 0)

	originator := NewOriginator()
	originator.SetState("state1")
	originator.SetState("state2")
	savedStatus = append(savedStatus, originator.CreateMemento())
	originator.SetState("state3")
	savedStatus = append(savedStatus, originator.CreateMemento())
	originator.SetState("state4")
	originator.Restore(savedStatus[1])

	// Output:
	// Originator: Setting state to  state1
	// Originator: Setting state to  state2
	// Originator: Saving to Memento.
	// Originator: Setting state to  state3
	// Originator: Saving to Memento.
	// Originator: Setting state to  state4
	// Originator : State after restoring from memento :  state3
}
