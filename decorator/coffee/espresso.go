package coffee

type Espresso struct{}

var _ Beverage = (*Espresso)(nil)

func (e *Espresso) Cost() float64 {
	return 1.99
}

func (e *Espresso) Description() string {
	return "Espresso"
}
