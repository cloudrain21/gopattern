package templatemethod

import "fmt"

type DefaultConnectAlgorithmOperator struct{}

func NewDefaultConnectAlgorithmOperator() ConnectAlgorithmOperator {
	return &DefaultConnectAlgorithmOperator{}
}

func (d *DefaultConnectAlgorithmOperator) Authentication(user string, pass string) {
	fmt.Println("default authentication")
	fmt.Println("added authentication policy")
}

func (d *DefaultConnectAlgorithmOperator) Authorization(user string, pass string) {
	fmt.Println("default authorization")
	fmt.Println("added authorization policy")
}
