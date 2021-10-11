package chain_of_responsibility

import "fmt"

type EvenHandler struct {
	name string
	next Handler
}

func NewEvenHandler(name string) *EvenHandler {
	return &EvenHandler{
		name: name,
		next: nil,
	}
}

func (h *EvenHandler) resolve(n int) bool {
	if n % 2 == 0 {
		return true
	} else {
		return false
	}
}

func (h *EvenHandler) done(n int) {
	fmt.Println(h.name, " resolved ", n)
}

func (h *EvenHandler) SetNext(n Handler) {
	h.next = n
}

// 상속을 지원한다면 이는 parent 로 올릴 수 있다.
func (h *EvenHandler)HandleRequest(n int) {
	if h.resolve(n) {
		h.done(n)
	} else {
		if h.next != nil {
			h.next.HandleRequest(n)
		} else {
			fmt.Println("no object resolved")
		}
	}
}
