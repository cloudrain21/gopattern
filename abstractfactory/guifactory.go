package abstractfactory

type GUIFactory interface {
	CreateButton() Button
	CreateScroll() Scroll
	CreatePopup() Popup
}

type Button interface {
	Click()
}

type Scroll interface {
	Up()
	Down()
}

type Popup interface {
	Pop()
}
