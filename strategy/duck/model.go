package duck

import "github.com/adambkaplan/go-design-patterns/strategy/behavior"

// ModelDuck represents a toy model duck, whose flying and quacking
// behavior can be changed. By default it can't fly and doesn't quack.
type ModelDuck struct {
	GenericDuck
}

var _ Duck = (*ModelDuck)(nil)

// Display prints a message to the configured writer.
func (m *ModelDuck) Display() (err error) {
	m.initWriter()
	_, err = m.Writer.WriteString("I'm a model duck\n")
	return
}

// PerformFly prints a flying action to the configured writer. If the flying behavior is not
// specified, it defaults to a "can't fly" action.
func (m *ModelDuck) PerformFly() error {
	m.initWriter()
	if m.FlyBehavior == nil {
		m.FlyBehavior = &behavior.FlyNoWay{}
	}
	return m.FlyBehavior.Fly(m.Writer)
}

// PerformQuack prints a quacking action to the configured writer. If the quacking behavior is not
// specified, it defaults to a "silent" quack.
func (m *ModelDuck) PerformQuack() error {
	m.initWriter()
	if m.QuackBehavior == nil {
		m.QuackBehavior = &behavior.MuteQuack{}
	}
	return m.QuackBehavior.Quack(m.Writer)
}
