package ny

import "github.com/adambkaplan/go-design-patterns/factory/pizza"

type PizzaStore struct {
	pizza.BasicStore
}

var _ pizza.Store = (*PizzaStore)(nil)
