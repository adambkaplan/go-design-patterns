package duck

import (
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

	// SetFlyBehavior sets the flying behavior of the duck.
	SetFlyBehavior(b behavior.Flyable)

	// SetQuackBehavior sets the quacking behavior of the duck.
	SetQuackBehavior(b behavior.Quackable)

	// Swim causes the duck to swim.
	Swim() error
}
