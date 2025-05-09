package coffee

type Soy struct {
	Condiment
}

func (s *Soy) Cost() float64 {
	return s.Condiment.Cost() + 0.15
}

func (s *Soy) Description() string {
	return s.Condiment.Description() + ", Soy"
}
