package factorymethod

const (
	METHOD1 = iota
	METHOD2
)

type ElevatorManager struct {
	sched Scheduler
}

func (e *ElevatorManager) RequestElevator() {
	e.sched.schedule()
}

func NewElevatorManager(method int) *ElevatorManager {
	elevatorMgr := new(ElevatorManager)

	switch method {
	case METHOD1:
		elevatorMgr.sched = NewMethod1Scheduler()
	case METHOD2:
		elevatorMgr.sched = NewMethod2Scheduler()
	}
	return elevatorMgr
}
