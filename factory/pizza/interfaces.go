package pizza

import "io"

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
type PizzaCreator func(writer io.StringWriter, pizzaType string) Pizza

type IngredientFactory interface {

	// CreateDough creates the permitted dough
	CreateDough() Dough

	// CreateSauce creates the permitted sauce
	CreateSauce() Sauce

	// CreateCheese creates the permitted cheese
	CreateCheese() Cheese

	// CreateVeggies creates the permitted veggies
	CreateVeggies() []Veggies

	// CreatePepperoni creates the permitted pepperoni
	CreatePepperoni() Pepperoni

	// CreateClams creates the permitted clams
	CreateClams() Clams
}
