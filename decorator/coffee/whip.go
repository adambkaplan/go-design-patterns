package coffee

type Whip struct {
	Condiment
}

func (w *Whip) Cost() float64 {
	return w.Condiment.Cost() + 0.1
}

func (w *Whip) Description() string {
	return w.Condiment.Description() + ", Whip"
}
