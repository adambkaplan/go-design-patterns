package behavior

import "io"

type Quackable interface {
	Quack(writer io.StringWriter)
}

type Quack struct{}

func (q *Quack) Quack(writer io.StringWriter) {
	writer.WriteString("Quack\n")
}

type MuteQuack struct{}

func (m *MuteQuack) Quack(writer io.StringWriter) {
	writer.WriteString("<< Silence >>\n")
}

type Squeak struct{}

var _ Quackable = (*Squeak)(nil)

// Quack squeaks.
func (s *Squeak) Quack(writer io.StringWriter) {
	writer.WriteString("Squeak")
}
