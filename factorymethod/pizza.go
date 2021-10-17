package factorymethod

type Pizza interface {
	GetPrice()
}

type PizzaType int

const (
	HamMushroomPizzaType PizzaType = iota
	DeluxePizzaType
	SeafoodPizzaType
)
