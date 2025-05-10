package ny

import "github.com/adambkaplan/go-design-patterns/factory/pizza"

type NYIngredients struct{}

// CreateCheese implements pizza.IngredientFactory.
func (i NYIngredients) CreateCheese() pizza.Cheese {
	return pizza.ShreddedMozarella
}

// CreateClams implements pizza.IngredientFactory.
func (i NYIngredients) CreateClams() pizza.Clams {
	return pizza.LIFreshClams
}

// CreateDough implements pizza.IngredientFactory.
func (i NYIngredients) CreateDough() pizza.Dough {
	return pizza.ThinCrust
}

// CreatePepperoni implements pizza.IngredientFactory.
func (i NYIngredients) CreatePepperoni() pizza.Pepperoni {
	return pizza.SlicedPepperoni
}

// CreateSauce implements pizza.IngredientFactory.
func (i NYIngredients) CreateSauce() pizza.Sauce {
	return pizza.Marinara
}

// CreateVeggies implements pizza.IngredientFactory.
func (i NYIngredients) CreateVeggies() []pizza.Veggies {
	return []pizza.Veggies{
		pizza.Garlic,
		pizza.Onion,
		pizza.Mushrooms,
		pizza.RedPepper,
	}
}

var _ pizza.IngredientFactory = (*NYIngredients)(nil)
