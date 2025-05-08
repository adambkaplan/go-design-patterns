package duck

import (
	"io"
	"os"

	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
)

// Duck describes all the actions a duck-like object can do.
type Duck interface {

	// Display shows the duck on the configured output.
	Display() error

	// PerformFly causes the duck to fly.
	PerformFly() error

	// PerformQuack causes the duck to make a sound.
	PerformQuack() error

	// Swim causes the duck to swim.
	Swim() error
}

// MallardDuck represents a real Mallard duck that flies and quacks.
type MallardDuck struct {

	// Writer is where the text output of this duck's actions are written.
	Writer io.StringWriter

	flyBehavior   behavior.Flyable
	quackBehavior behavior.Quackable
}

var _ Duck = (*MallardDuck)(nil)

func (m *MallardDuck) initWriter() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
}

// Display prints a message to the configured writer.
func (m *MallardDuck) Display() (err error) {
	m.initWriter()
	_, err = m.Writer.WriteString("I'm a real Mallard duck\n")
	return
}

// PerformFly prints a flying action to the configured writer.
func (m *MallardDuck) PerformFly() error {
	m.initWriter()
	if m.flyBehavior == nil {
		m.flyBehavior = &behavior.FlyWithWings{}
	}
	return m.flyBehavior.Fly(m.Writer)
}

// PerformQuack prints a quacking action to the configured writer.
func (m *MallardDuck) PerformQuack() error {
	m.initWriter()
	if m.quackBehavior == nil {
		m.quackBehavior = &behavior.Quack{}
	}
	return m.quackBehavior.Quack(m.Writer)
}

// Swim prints a swimming action to the configured writer.
func (m *MallardDuck) Swim() (err error) {
	m.initWriter()
	_, err = m.Writer.WriteString("All ducks float, even decoys!\n")
	return
}

// ModelDuck represents a toy model duck, whose flying and quacking
// behavior can be changed. By default it can't fly and doesn't quack.
type ModelDuck struct {

	// FlyBehavior is the flying behavior of this duck.
	FlyBehavior behavior.Flyable

	// QuackBehavior is the quacking behavior of this duck.
	QuackBehavior behavior.Quackable

	// Writer is where the text output of this duck's actions are written.
	Writer io.StringWriter
}

var _ Duck = (*ModelDuck)(nil)

func (m *ModelDuck) initWriter() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
}

// Display prints a message to the configured writer.
func (m *ModelDuck) Display() (err error) {
	m.initWriter()
	_, err = m.Writer.WriteString("I'm a model duck\n")
	return
}

// PerformFly prints a flying action to the configured writer.
func (m *ModelDuck) PerformFly() error {
	m.initWriter()
	if m.FlyBehavior == nil {
		m.FlyBehavior = &behavior.FlyNoWay{}
	}
	return m.FlyBehavior.Fly(m.Writer)
}

// PerformQuack prints a quacking action to the configured writer.
func (m *ModelDuck) PerformQuack() error {
	m.initWriter()
	if m.QuackBehavior == nil {
		m.QuackBehavior = &behavior.MuteQuack{}
	}
	return m.QuackBehavior.Quack(m.Writer)
}

// Swim prints a swimming action to the configured writer.
func (m *ModelDuck) Swim() (err error) {
	m.initWriter()
	_, err = m.Writer.WriteString("All ducks float, even decoys!\n")
	return
}
