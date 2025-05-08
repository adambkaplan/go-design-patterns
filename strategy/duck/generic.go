package duck

import (
	"io"
	"os"

	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
)

// GenericDuck represents a non-specific duck. By default it can fly and quack, but these behaviors
// can be changed.
type GenericDuck struct {

	// FlyBehavior is the flying behavior of this duck.
	FlyBehavior behavior.Flyable

	// QuackBehavior is the quacking behavior of this duck.
	QuackBehavior behavior.Quackable

	// Writer is where the text output of this duck's actions are written.
	Writer io.StringWriter
}

var _ Duck = (*GenericDuck)(nil)

// initWriter initializes the writer to `os.Stdout` if it has not been set.
func (g *GenericDuck) initWriter() {
	if g.Writer == nil {
		g.Writer = os.Stdout
	}
}

// Display displays the duck to the writer.
func (g *GenericDuck) Display() (err error) {
	g.initWriter()
	_, err = g.Writer.WriteString("I'm a non-descript duck\n")
	return
}

// PerformFly prints a flying action to the configured writer. If the flying behavior is not set,
// it defaults to flying with wings.
func (g *GenericDuck) PerformFly() error {
	g.initWriter()
	if g.FlyBehavior == nil {
		g.FlyBehavior = &behavior.FlyWithWings{}
	}
	return g.FlyBehavior.Fly(g.Writer)
}

// PerformQuack prints a quacking action to the configured writer. If the quacking behavior is not
// set, it defaults to a typical duck quack.
func (g *GenericDuck) PerformQuack() error {
	g.initWriter()
	if g.QuackBehavior == nil {
		g.QuackBehavior = &behavior.Quack{}
	}
	return g.QuackBehavior.Quack(g.Writer)
}

// SetFlyBehavior sets the flying behavior for the duck.
func (g *GenericDuck) SetFlyBehavior(b behavior.Flyable) {
	g.FlyBehavior = b
}

// SetQuackBehavior sets the quacking behavior for the duck.
func (g *GenericDuck) SetQuackBehavior(b behavior.Quackable) {
	g.QuackBehavior = b
}

// Swim prints a swimming action to the configured writer.
func (g *GenericDuck) Swim() (err error) {
	g.initWriter()
	_, err = g.Writer.WriteString("All ducks float, even decoys!\n")
	return
}
