package abstractfactory

func ExampleBuildGUIFactory() {
	factory := NewWindowsGUIFactory()

	button := factory.CreateButton()
	scroll := factory.CreateScroll()
	popup := factory.CreatePopup()

	button.Click()
	scroll.Up()
	scroll.Down()
	popup.Pop()

	factory = NewLinuxGUIFactory()
	button = factory.CreateButton()
	scroll = factory.CreateScroll()
	popup = factory.CreatePopup()

	button.Click()
	scroll.Up()
	scroll.Down()
	popup.Pop()

	// Output:
	// WindowsButton Click
	// WindowsScroll Up
	// WindowsScroll Down
	// WindowsPopup Pop
	// LinuxButton Click
	// LinuxScroll Up
	// LinuxScroll Down
	// LinuxPopup Pop
}
