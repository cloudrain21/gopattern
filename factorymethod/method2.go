package factorymethod

import "fmt"

func NewMethod2Scheduler() *Method2Scheduler {
	return &Method2Scheduler{}
}

type Method2Scheduler struct {
}

func (m *Method2Scheduler)schedule() {
	fmt.Println("method2 scheduling")
}