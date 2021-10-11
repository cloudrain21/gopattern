package chain_of_responsibility

import "fmt"

type OddHandler struct {
	name string
	next Handler
}

func NewOddHandler(name string) *OddHandler {
	return &OddHandler{
		name: name,
		next: nil,
	}
}

func (h *OddHandler) resolve(n int) bool {
	if n % 2 == 0 {
		return false
	} else {
		return true
	}
}

func (h *OddHandler) done(n int) {
	fmt.Println(h.name, " resolved ", n)
}

func (h *OddHandler) SetNext(n Handler) {
	h.next = n
}

// 상속을 지원한다면 이는 parent 로 올릴 수 있다.
func (h *OddHandler)HandleRequest(n int) {
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