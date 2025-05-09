package factory

import (
	"fmt"

	"github.com/adambkaplan/go-design-patterns/factory/pizza"
	"github.com/adambkaplan/go-design-patterns/factory/pizza/chicago"
	"github.com/adambkaplan/go-design-patterns/factory/pizza/ny"
)

func RunPizzaStores() {

	nyStore := &ny.PizzaStore{}
	chicagoStore := &chicago.PizzaStore{}

	personStore := map[string]pizza.Store{
		"Ethan": nyStore,
		"Joel":  chicagoStore,
	}
	pizzaTypes := []string{"cheese", "clam", "pepperoni", "veggie"}

	for _, pizzaType := range pizzaTypes {
		for person, store := range personStore {
			orderPizza(person, pizzaType, store)
		}
	}
}

func orderPizza(person string, pizzaType string, store pizza.Store) {
	pizza, err := store.OrderPizza(pizzaType)
	if err != nil {
		fmt.Printf("Uh oh %s, something went wrong! %v\n", person, err)
		return
	}
	fmt.Printf("%s ordered a %s\n", person, pizza.GetName())
}
