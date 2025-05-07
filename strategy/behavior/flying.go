package behavior

// Flyable is the behavior of flying things.
type Flyable interface {

	// Fly executes the flying behavior.
	Fly()
}

type FlyRocketPowered struct{}

var _ Flyable = (*FlyRocketPowered)(nil)

func (f *FlyRocketPowered) Fly() {
	panic("unimplemented")
}
