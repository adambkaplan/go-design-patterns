package coffee

type Beverage interface {
	// Note that for real financial transactions, you should not use floating point numbers!

	// Cost is the cost of the beverage.
	Cost() float64

	// Decription describes the beverage.
	Description() string
}

// unknownBeverage is an empty instance of the Beverage interface. Its purpose is to ensure empty
// Condiment instances do not panic with nil pointer dereferences.
type unknownBeverage struct{}

func (u *unknownBeverage) Cost() float64 {
	return 0.0
}

func (u *unknownBeverage) Description() string {
	return "unknown"
}
