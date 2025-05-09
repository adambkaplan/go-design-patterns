package coffee

type Mocha struct {
	Condiment
}

func (m *Mocha) Cost() float64 {
	return m.Condiment.Cost() + 0.20
}

func (m *Mocha) Description() string {
	return m.Condiment.Description() + ", Mocha"
}
