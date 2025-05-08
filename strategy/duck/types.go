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

func (m *MallardDuck) initWriter() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
}

// Display implements Duck.
func (m *MallardDuck) Display() {
	m.initWriter()
	m.Writer.WriteString("I'm a real Mallard duck\n")
}

// PerformFly implements Duck.
func (m *MallardDuck) PerformFly() {
	m.initWriter()
	if m.flyBehavior == nil {
		m.flyBehavior = &behavior.FlyWithWings{}
	}
	m.flyBehavior.Fly(m.Writer)
}

// PerformQuack implements Duck.
func (m *MallardDuck) PerformQuack() {
	m.initWriter()
	if m.quackBehavior == nil {
		m.quackBehavior = &behavior.Quack{}
	}
	m.quackBehavior.Quack(m.Writer)
}

// Swim implements Duck.
func (m *MallardDuck) Swim() {
	m.initWriter()
	m.Writer.WriteString("All ducks float, even decoys!\n")
}

type ModelDuck struct {
	FlyBehavior   behavior.Flyable
	QuackBehavior behavior.Quackable
	Writer        io.StringWriter
}

var _ Duck = (*ModelDuck)(nil)

func (m *ModelDuck) initWriter() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
}

// Display implements Duck.
func (m *ModelDuck) Display() {
	m.initWriter()
	m.Writer.WriteString("I'm a model duck\n")
}

// PerformFly implements Duck.
func (m *ModelDuck) PerformFly() {
	m.initWriter()
	if m.FlyBehavior == nil {
		m.FlyBehavior = &behavior.FlyNoWay{}
	}
	m.FlyBehavior.Fly(m.Writer)
}

// PerformQuack implements Duck.
func (m *ModelDuck) PerformQuack() {
	m.initWriter()
	if m.QuackBehavior == nil {
		m.QuackBehavior = &behavior.MuteQuack{}
	}
	m.QuackBehavior.Quack(m.Writer)
}

// Swim implements Duck.
func (m *ModelDuck) Swim() {
	m.initWriter()
	m.Writer.WriteString("All ducks float, even decoys!\n")
}
