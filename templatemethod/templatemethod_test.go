package templatemethod

func ExampleConnectHelper() {
	// default connection helper
	connectHelper := NewAbstConnectHelper()
	connectHelper.Connect("user", "pass")

	// other connection helper
	otherConnectionHelper := NewOtherConnectAlgorithmOperator()
	connectHelper.ChangeConnectionAlgorithmOperator(otherConnectionHelper)
	connectHelper.Connect("user", "pass")

	// Output:
	// default authentication
	// added authentication policy
	// default authorization
	// added authorization policy
	// other authentication
	// added other authentication policy
	// other authorization
	// added other authorization policy

}
