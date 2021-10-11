package observer

import "fmt"

type ConcreteObserver1 struct {

}

func NewConcreteObserver1() *ConcreteObserver1 {
	return &ConcreteObserver1{}
}

func (o *ConcreteObserver1)Notify(i int) {
	fmt.Println("concrete observer 1 : ", i)
}
