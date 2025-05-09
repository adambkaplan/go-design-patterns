package ny

import "github.com/adambkaplan/go-design-patterns/factory/pizza"

type PizzaStore struct {
	pizza.BasicStore
}

// OrderPizza implements pizza.Store.
func (p *PizzaStore) OrderPizza(pizzaType string) (pizza.Pizza, error) {
	return &pizza.BasicPizza{}, nil
}

var _ pizza.Store = (*PizzaStore)(nil)
