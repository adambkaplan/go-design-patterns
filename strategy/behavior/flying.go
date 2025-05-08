package behavior

import "io"

// Flyable is the behavior of flying things.
type Flyable interface {

	// Fly writes the flying behavior to the provided writer.
	Fly(writer io.StringWriter) error
}

// FlyWithWings is a flying action with wings.
type FlyWithWings struct{}

// Fly writes the flying behavior to the provided writer.
func (f *FlyWithWings) Fly(writer io.StringWriter) (err error) {
	_, err = writer.WriteString("I'm flying!!\n")
	return
}

// FlyNoWay is a "can't fly" action.
type FlyNoWay struct{}

// Fly writes the flying behavior to the provided writer.
func (f *FlyNoWay) Fly(writer io.StringWriter) (err error) {
	_, err = writer.WriteString("I can't fly\n")
	return
}

// FlyRocketPowered is a flying action assisted by rockets.
type FlyRocketPowered struct{}

var _ Flyable = (*FlyRocketPowered)(nil)

// Fly writes the flying behavior to the provided writer.
func (f *FlyRocketPowered) Fly(writer io.StringWriter) (err error) {
	_, err = writer.WriteString("I'm flying with a rocket!\n")
	return
}
