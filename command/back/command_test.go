package command

import "fmt"

// 함수 자체를 slice 에 넣어두고 꺼내서 수행하는 방식
func ExampleCommand1() {
	l := []func(){}

	l = append(l, func() {
		fmt.Println(111)
	})
	l = append(l, func() {
		fmt.Println(222)
	})

	for _, v := range l {
		v()
	}

	// Output:
	// 111
	// 222
}


// 특정 타입을 Command interface 형으로 slice 로 넣어두고 꺼내서 수행하는 방식
func ExampleCommand2() {
	l := []Command{}

	l = append(l, NewStringCommand("aaa"))
	l = append(l, NewStringCommand("bbb"))
	l = append(l, NewStringCommand("ccc"))

	for _, v := range l {
		v.Execute()
	}

	// Output:
	// aaa
	// bbb
	// ccc
}


// 다양한 명령 객체를 Command interface slice 에 넣어두고, 이를 같은 interface 로 꺼내서 수행하는 방식
func ExampleCommand3() {
	l := []Command{}

	l = append(l, NewIntCommand(111))
	l = append(l, NewStringCommand("aaa"))
	l = append(l, NewIntCommand(222))
	l = append(l, NewStringCommand("bbb"))

	for _, v := range l {
		v.Execute()
	}

	// Output :
	// 111
	// aaa
	// 222
	// bbb
}