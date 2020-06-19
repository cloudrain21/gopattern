package factorymethod

func ExampleElevatorManager() {
	elevatorMgr := NewElevatorManager(METHOD1)
	elevatorMgr.RequestElevator()
	elevatorMgr = NewElevatorManager(METHOD2)
	elevatorMgr.RequestElevator()

	// Output:
	// method1 scheduling
	// method2 scheduling
}
