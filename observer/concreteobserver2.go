package observer

import "fmt"

type ConcreteObserver2 struct {

}

func NewConcreteObserver2() *ConcreteObserver2 {
	return &ConcreteObserver2{}
}

func (o *ConcreteObserver2)Notify(i int) {
	fmt.Println("concrete observer 2 : ", i)
}
