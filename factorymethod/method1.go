package factorymethod

import "fmt"

func NewMethod1Scheduler() *Method1Scheduler {
	return &Method1Scheduler{}
}

type Method1Scheduler struct {
}

func (m *Method1Scheduler) schedule() {
	fmt.Println("method1 scheduling")
}
