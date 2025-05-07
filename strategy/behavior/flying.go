package behavior

import "io"

// Flyable is the behavior of flying things.
type Flyable interface {

	// Fly executes the flying behavior.
	Fly(writer io.StringWriter)
}

type FlyWithWings struct{}

func (f *FlyWithWings) Fly(writer io.StringWriter) {
	writer.WriteString("I'm flying!!\n")
}

type FlyNoWay struct{}

func (f *FlyNoWay) Fly(writer io.StringWriter) {
	writer.WriteString("I can't fly\n")
}

type FlyRocketPowered struct{}

var _ Flyable = (*FlyRocketPowered)(nil)

func (f *FlyRocketPowered) Fly(writer io.StringWriter) {
	writer.WriteString("I'm flying with a rocket!")
}
