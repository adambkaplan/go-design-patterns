package duck

import (
	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
)

type Duck interface {
	Display()

	PerformFly()

	PerformQuack()

	Swim()
}

type MallardDuck struct {
}

var _ Duck = (*MallardDuck)(nil)

// Display implements Duck.
func (m *MallardDuck) Display() {
	panic("unimplemented")
}

// PerformFly implements Duck.
func (m *MallardDuck) PerformFly() {
	panic("unimplemented")
}

// PerformQuack implements Duck.
func (m *MallardDuck) PerformQuack() {
	panic("unimplemented")
}

// Swim implements Duck.
func (m *MallardDuck) Swim() {
	panic("unimplemented")
}

type ModelDuck struct {
	FlyBehavior behavior.Flyable
}

var _ Duck = (*ModelDuck)(nil)

// Display implements Duck.
func (m *ModelDuck) Display() {
	panic("unimplemented")
}

// PerformFly implements Duck.
func (m *ModelDuck) PerformFly() {
	panic("unimplemented")
}

// PerformQuack implements Duck.
func (m *ModelDuck) PerformQuack() {
	panic("unimplemented")
}

// Swim implements Duck.
func (m *ModelDuck) Swim() {
	panic("unimplemented")
}
