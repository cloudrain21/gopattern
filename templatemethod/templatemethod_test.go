package templatemethod

func ExampleConnectHelper() {
	ConnectHelper := NewAbstConnectHelper()
	ConnectHelper.Connect("user", "pass")

	// Output:
	// default authentication
	// default authorization
}