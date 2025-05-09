package pizza

import (
	"io"
	"os"
)

type BasicPizza struct {
	Writer io.StringWriter
	Name   string

	Dough     Dough
	Sauce     Sauce
	Cheese    Cheese
	Veggies   []Veggies
	Pepperoni Pepperoni
	Clams     Clams
}

var _ (Pizza) = (*BasicPizza)(nil)

func (b *BasicPizza) InitWriter() {
	if b.Writer == nil {
		b.Writer = os.Stdout
	}
}

// Bake implements Pizza.
func (b *BasicPizza) Bake() (err error) {
	b.InitWriter()
	_, err = b.Writer.WriteString("Bake for 25 minutes at 350\n")
	return
}

// Box implements Pizza.
func (b *BasicPizza) Box() (err error) {
	b.InitWriter()
	_, err = b.Writer.WriteString("Place pizza in official PizzaStore box\n")
	return
}

// Cut implements Pizza.
func (b *BasicPizza) Cut() (err error) {
	b.InitWriter()
	_, err = b.Writer.WriteString("Cutting the pizza into diagonal slices\n")
	return
}

func (b *BasicPizza) Prepare() error {
	return nil
}

// String implements Pizza.
func (b *BasicPizza) String() string {
	return b.GetName()
}

func (b *BasicPizza) GetName() string {
	if len(b.Name) == 0 {
		b.Name = "Basic Pizza"
	}
	return b.Name
}

// CreateBasicPizza always creates a basic pizza, regardless of input pizza type.
func CreateBasicPizza(writer io.StringWriter, pizzaType string) Pizza {
	return &BasicPizza{
		Writer: writer,
	}
}
