package ny

import (
	"io"

	"github.com/adambkaplan/go-design-patterns/factory/pizza"
)

func NewPizzaStore() pizza.Store {
	return &pizza.BasicStore{
		CreatePizza: createNYPizza,
	}
}

func createNYPizza(writer io.StringWriter, pizzaType string) (newPizza pizza.Pizza) {

	ingredients := NYIngredients{}

	switch pizzaType {
	case "cheese":
		newPizza = &pizza.CheesePizza{
			BasicPizza: pizza.BasicPizza{
				Name: "New York Style Cheese Pizza",
			},
			Ingredients: ingredients,
		}
	}
	return
}
