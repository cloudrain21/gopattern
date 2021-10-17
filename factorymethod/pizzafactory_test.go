package factorymethod

func ExamplePizzaFactory() {
	pizzaFactory := NewPizzaFactory()
	pizzaFactory.CreatePizza(HamMushroomPizzaType).GetPrice()
	pizzaFactory.CreatePizza(DeluxePizzaType).GetPrice()
	pizzaFactory.CreatePizza(SeafoodPizzaType).GetPrice()

	// Output:
	// ham and mushroom : 100
	// deluxe : 200
	// seafood : 300
}
