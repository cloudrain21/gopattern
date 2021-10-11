package chain_of_responsibility

func ExampleHandler() {
	var oddHandler Handler
	var evenHandler Handler

	oddHandler = NewOddHandler("odd_handler")
	evenHandler = NewEvenHandler("even_handler")

	oddHandler.SetNext(evenHandler)

	for i := 0; i < 10; i++ {
		oddHandler.HandleRequest(i)
	}

	// Output:
	// even_handler  resolved  0
	// odd_handler  resolved  1
	// even_handler  resolved  2
	// odd_handler  resolved  3
	// even_handler  resolved  4
	// odd_handler  resolved  5
	// even_handler  resolved  6
	// odd_handler  resolved  7
	// even_handler  resolved  8
	// odd_handler  resolved  9
}
