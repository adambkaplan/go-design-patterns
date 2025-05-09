package pizza

type BasicStore struct {
	CreatePizza PizzaCreator
}

var _ Store = (*BasicStore)(nil)

// OrderPizza implements Store.
func (b *BasicStore) OrderPizza(pizzaType string) (Pizza, error) {
	if b.CreatePizza == nil {
		b.CreatePizza = CreateBasicPizza
	}
	pizza := b.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Box()
	pizza.Cut()
	return pizza, nil
}
