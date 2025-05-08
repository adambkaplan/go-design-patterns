package behavior

import "io"

// Quackable represents any object capable of quacking.
type Quackable interface {

	// Quack writes a quacking action to the provided writer.
	Quack(writer io.StringWriter) error
}

// Quack is a typical duck quack action.
type Quack struct{}

// Quack writes a duck quack to the provided writer.
func (q *Quack) Quack(writer io.StringWriter) (err error) {
	_, err = writer.WriteString("Quack\n")
	return
}

// MuteQuack is a "silent" quack action.
type MuteQuack struct{}

// Quack writes a "silent" quack to the provided writer.
func (m *MuteQuack) Quack(writer io.StringWriter) (err error) {
	_, err = writer.WriteString("<< Silence >>\n")
	return
}

// Squeak is a quack that sounds like a "squeak", as in the sound a
// rubber duck makes.
type Squeak struct{}

var _ Quackable = (*Squeak)(nil)

// Quack writes a "squeaking" quack to the provided writer.
func (s *Squeak) Quack(writer io.StringWriter) (err error) {
	_, err = writer.WriteString("Squeak\n")
	return
}
