package chicago

import (
	"io"

	"github.com/adambkaplan/go-design-patterns/factory/pizza"
)

func NewPizzaStore() pizza.Store {
	return &pizza.BasicStore{
		CreatePizza: createChicagoPizza,
	}
}
func createChicagoPizza(writer io.StringWriter, pizzaType string) (newPizza pizza.Pizza) {

	ingredients := ChicagoIngredients{}

	switch pizzaType {
	case "cheese":
		newPizza = &pizza.CheesePizza{
			BasicPizza: pizza.BasicPizza{
				Name:   "Chicago Style Deep Dish Cheese Pizza",
				Writer: writer,
			},
			Ingredients: ingredients,
		}
	}
	return
}
