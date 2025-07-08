package chicago

import "github.com/adambkaplan/go-design-patterns/factory/pizza"

type ChicagoIngredients struct{}

// CreateCheese implements pizza.IngredientFactory.
func (i ChicagoIngredients) CreateCheese() pizza.Cheese {
	return pizza.ShreddedMozarella
}

// CreateClams implements pizza.IngredientFactory.
func (i ChicagoIngredients) CreateClams() pizza.Clams {
	return pizza.FrozenClams
}

// CreateDough implements pizza.IngredientFactory.
func (i ChicagoIngredients) CreateDough() pizza.Dough {
	return pizza.ThickCrust
}

// CreatePepperoni implements pizza.IngredientFactory.
func (i ChicagoIngredients) CreatePepperoni() pizza.Pepperoni {
	return pizza.SlicedPepperoni
}

// CreateSauce implements pizza.IngredientFactory.
func (i ChicagoIngredients) CreateSauce() pizza.Sauce {
	return pizza.PlumTomato
}

// CreateVeggies implements pizza.IngredientFactory.
func (i ChicagoIngredients) CreateVeggies() []pizza.Veggies {
	return []pizza.Veggies{
		pizza.BlackOlives,
		pizza.Spinach,
		pizza.Eggplant,
	}
}

var _ pizza.IngredientFactory = (*ChicagoIngredients)(nil)
