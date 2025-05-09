package coffee

// Condiment is a base decorator that other beverage decorators embed.
type Condiment struct {
	Beverage Beverage
}

var _ Beverage = (*Condiment)(nil)

// Cost computes the cost of the beverage with the condiment.
func (c *Condiment) Cost() float64 {
	if c.Beverage == nil {
		c.Beverage = &unknownBeverage{}
	}
	return c.Beverage.Cost()
}

// Description describes the beverage with the condiment.
func (c *Condiment) Description() string {
	if c.Beverage == nil {
		c.Beverage = &unknownBeverage{}
	}
	return c.Beverage.Description()
}
