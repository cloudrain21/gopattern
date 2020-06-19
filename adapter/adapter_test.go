package adapter

import "fmt"

func ExampleMathAdapter() {
	fmt.Println(Twice(10))
	fmt.Println(Half(10))

	adapter := NewStringMathAdapter()
	fmt.Println(adapter.Twice("10"))
	fmt.Println(adapter.Half("10"))

	// Output:
	// 20
	// 5
	// Adapter : 20
	// Adapter : 5
}
