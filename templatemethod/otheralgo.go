package templatemethod

import "fmt"

type OtherConnectAlgorithmOperator struct{}

func NewOtherConnectAlgorithmOperator() ConnectAlgorithmOperator {
	return &OtherConnectAlgorithmOperator{}
}

func (d *OtherConnectAlgorithmOperator) Authentication(user string, pass string) {
	fmt.Println("other authentication")
	fmt.Println("added other authentication policy")
}

func (d *OtherConnectAlgorithmOperator) Authorization(user string, pass string) {
	fmt.Println("other authorization")
	fmt.Println("added other authorization policy")
}
