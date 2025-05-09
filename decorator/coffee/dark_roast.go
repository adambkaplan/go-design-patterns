package coffee

type DarkRoast struct{}

var _ Beverage = (*DarkRoast)(nil)

func (d *DarkRoast) Cost() float64 {
	return 0.99
}

func (d *DarkRoast) Description() string {
	return "Dark Roast Coffee"
}
