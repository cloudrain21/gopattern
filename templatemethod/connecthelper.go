package templatemethod

type ConnectHelper interface {
	Connect(string, string) error
}

// go 에는 상속이 없어서 AbstConnectHelper 에서 또 다른 interface 의
// operator 를 통해 Connect 를 구현하도록 하였다.
// 상속을 이용하지 않은 domain algorithm 의 구현이라고 보면 되겠다.
type ConnectAlgorithmOperator interface {
	Authentication(string, string)
	Authorization(string, string)
}

type AbstConnectHelper struct {
	algo ConnectAlgorithmOperator
}

func (a *AbstConnectHelper) Connect(user string, pass string) error {
	a.algo.Authentication(user, pass)
	a.algo.Authorization(user, pass)
	return nil
}

func NewAbstConnectHelper() *AbstConnectHelper {
	return &AbstConnectHelper{NewDefaultConnectAlgorithmOperator()}
}

func (a *AbstConnectHelper) ChangeConnectionAlgorithmOperator(algo ConnectAlgorithmOperator) {
	a.algo = algo
}
