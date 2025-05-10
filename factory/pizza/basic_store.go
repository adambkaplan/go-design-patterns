package pizza

import (
	"io"
	"os"
)

type BasicStore struct {
	Writer      io.StringWriter
	CreatePizza PizzaCreator
}

var _ Store = (*BasicStore)(nil)

// OrderPizza implements Store.
func (b *BasicStore) OrderPizza(pizzaType string) (Pizza, error) {
	if b.CreatePizza == nil {
		b.CreatePizza = CreateBasicPizza
	}
	if b.Writer == nil {
		b.Writer = os.Stdout
	}
	pizza := b.CreatePizza(b.Writer, pizzaType)

	if err := pizza.Prepare(); err != nil {
		return nil, err
	}
	if err := pizza.Bake(); err != nil {
		return nil, err
	}
	if err := pizza.Box(); err != nil {
		return nil, err
	}
	if err := pizza.Cut(); err != nil {
		return nil, err
	}
	return pizza, nil
}
