package factorymethod

type PizzaFactory struct {
}

func NewPizzaFactory() *PizzaFactory {
	return &PizzaFactory{}
}

func (f *PizzaFactory) CreatePizza(pizzaType PizzaType) Pizza {
	if pizzaType == HamMushroomPizzaType {
		return NewHamAndMushroom()
	} else if pizzaType == DeluxePizzaType {
		return NewDeluxe()
	} else {
		return NewSeafood()
	}
}
