package pizza

type Pizza interface {

	// GetName returns the name of the pizza.
	GetName() string

	// Bake bakes the pizza.
	Bake() error

	// Box boxes the pizza.
	Box() error

	// Cut cuts the pizza.
	Cut() error

	// Prepare prepares the pizza.
	Prepare() error

	// String represents the entire Pizza.
	String() string
}

type Store interface {

	// OrderPizza orders a pizza of the given type.
	OrderPizza(pizzaType string) (Pizza, error)
}

// PizzaCreator creates a pizza of the given type.
type PizzaCreator func(pizzaType string) Pizza
