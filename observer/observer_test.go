package observer

func ExampleObserver() {
	concreteObserver1 := NewConcreteObserver1()
	concreteObserver2 := NewConcreteObserver2()

	subject := NewSubject()

	subject.RegisterObserver(concreteObserver1)
	subject.RegisterObserver(concreteObserver2)

	subject.Run()

	// Output:
	// concrete observer 1 :  0
	// concrete observer 2 :  0
	// concrete observer 1 :  1
	// concrete observer 2 :  1
}
