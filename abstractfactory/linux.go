package abstractfactory

import "fmt"

type LinuxGUIFactory struct{}

type LinuxButton struct{}

type LinuxScroll struct{}

type LinuxPopup struct{}

func NewLinuxGUIFactory() GUIFactory {
	return &LinuxGUIFactory{}
}

func (w *LinuxGUIFactory) CreateButton() Button {
	return &LinuxButton{}
}

func (w *LinuxGUIFactory) CreateScroll() Scroll {
	return &LinuxScroll{}
}

func (w *LinuxGUIFactory) CreatePopup() Popup {
	return &LinuxPopup{}
}

func (b *LinuxButton) Click() {
	fmt.Println("LinuxButton Click")
}

func (b *LinuxScroll) Up() {
	fmt.Println("LinuxScroll Up")
}

func (b *LinuxScroll) Down() {
	fmt.Println("LinuxScroll Down")
}

func (b *LinuxPopup) Pop() {
	fmt.Println("LinuxPopup Pop")
}
