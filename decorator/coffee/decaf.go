package coffee

type Decaf struct{}

var _ Beverage = (*Decaf)(nil)

func (d *Decaf) Cost() float64 {
	return 1.05
}

func (d *Decaf) Description() string {
	return "Decaf Coffee"
}
