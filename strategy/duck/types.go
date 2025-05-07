package duck

import (
	"io"
	"os"

	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
)

type Duck interface {
	Display()

	PerformFly()

	PerformQuack()

	Swim()
}

type MallardDuck struct {
	Writer io.StringWriter

	flyBehavior   behavior.Flyable
	quackBehavior behavior.Quackable
}

var _ Duck = (*MallardDuck)(nil)

// Display implements Duck.
func (m *MallardDuck) Display() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
	m.Writer.WriteString("I am a real Mallard duck\n")
}

// PerformFly implements Duck.
func (m *MallardDuck) PerformFly() {
	if m.flyBehavior == nil {
		m.flyBehavior = &behavior.FlyWithWings{}
	}
	m.flyBehavior.Fly(m.Writer)
}

// PerformQuack implements Duck.
func (m *MallardDuck) PerformQuack() {
	if m.quackBehavior == nil {
		m.quackBehavior = &behavior.Quack{}
	}
	m.quackBehavior.Quack(m.Writer)
}

// Swim implements Duck.
func (m *MallardDuck) Swim() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
	m.Writer.WriteString("All ducks can swim, even decoys!\n")
}

type ModelDuck struct {
	FlyBehavior   behavior.Flyable
	QuackBehavior behavior.Quackable
	Writer        io.StringWriter
}

var _ Duck = (*ModelDuck)(nil)

// Display implements Duck.
func (m *ModelDuck) Display() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
	m.Writer.WriteString("I am a model duck\n")
}

// PerformFly implements Duck.
func (m *ModelDuck) PerformFly() {
	if m.FlyBehavior == nil {
		m.FlyBehavior = &behavior.FlyNoWay{}
	}
	m.FlyBehavior.Fly(m.Writer)
}

// PerformQuack implements Duck.
func (m *ModelDuck) PerformQuack() {
	if m.QuackBehavior == nil {
		m.QuackBehavior = &behavior.MuteQuack{}
	}
	m.QuackBehavior.Quack(m.Writer)
}

// Swim implements Duck.
func (m *ModelDuck) Swim() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
	m.Writer.WriteString("All ducks can swim, even decoys!\n")
}
