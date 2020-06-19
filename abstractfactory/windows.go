package abstractfactory

import "fmt"

type WindowsGUIFactory struct {}

type WindowsButton struct {}

type WindowsScroll struct {}

type WindowsPopup struct {}

func NewWindowsGUIFactory() GUIFactory {
	return &WindowsGUIFactory{}
}

func (w *WindowsGUIFactory)CreateButton() Button {
	return &WindowsButton{}
}

func (w *WindowsGUIFactory)CreateScroll() Scroll {
	return &WindowsScroll{}
}

func (w *WindowsGUIFactory)CreatePopup() Popup {
	return &WindowsPopup{}
}

func (w *WindowsButton)Click() {
	fmt.Println("WindowsButton Click")
}

func (w *WindowsScroll)Up() {
	fmt.Println("WindowsScroll Up")
}

func (w *WindowsScroll)Down() {
	fmt.Println("WindowsScroll Down")
}

func (w *WindowsPopup)Pop() {
	fmt.Println("WindowsPopup Pop")
}