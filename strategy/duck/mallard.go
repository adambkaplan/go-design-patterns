package duck

// MallardDuck represents a real Mallard duck that flies and quacks.
type MallardDuck struct {
	GenericDuck
}

// Display prints a message to the configured writer.
func (m *MallardDuck) Display() (err error) {
	m.initWriter()
	_, err = m.Writer.WriteString("I'm a real Mallard duck\n")
	return
}
