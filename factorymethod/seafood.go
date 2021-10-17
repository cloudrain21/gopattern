package factorymethod

import "fmt"

type Seafood struct {
}

func NewSeafood() *Seafood {
	return &Seafood{}
}

func (m *Seafood) GetPrice() {
	fmt.Println("seafood : 300")
}
