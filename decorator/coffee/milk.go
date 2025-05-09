package coffee

type Milk struct {
	Condiment
}

func (m *Milk) Cost() float64 {
	return m.Condiment.Cost() + 0.10
}

func (m *Milk) Description() string {
	return m.Condiment.Description() + ", Milk"
}
