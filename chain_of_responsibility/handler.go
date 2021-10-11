package chain_of_responsibility

// https://lktprogrammer.tistory.com/45
type Handler interface {
	HandleRequest(n int)
	SetNext(n Handler)
	resolve(n int) bool
	done(n int)
}
