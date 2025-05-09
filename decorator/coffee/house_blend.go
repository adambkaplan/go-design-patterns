package coffee

type HouseBlend struct{}

var _ Beverage = (*HouseBlend)(nil)

func (h *HouseBlend) Cost() float64 {
	return 0.89
}

func (h *HouseBlend) Description() string {
	return "House Blend Coffee"
}
