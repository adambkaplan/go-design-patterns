package pizza

import (
	"fmt"
	"strings"
)

type CheesePizza struct {
	BasicPizza

	Ingredients IngredientFactory
}

func (c *CheesePizza) Prepare() error {
	if c.Ingredients == nil {
		return fmt.Errorf("can't prepare %s: no ingredients", c.Name)
	}

	outBuilder := &strings.Builder{}
	fmt.Fprintf(outBuilder, "Preparing %s\n", c.Name)

	c.Dough = c.Ingredients.CreateDough()
	fmt.Fprintf(outBuilder, "Tossing %s...\n", c.Dough)

	c.Sauce = c.Ingredients.CreateSauce()
	fmt.Fprintf(outBuilder, "Adding %s...\n", c.Sauce)

	fmt.Fprintf(outBuilder, "Adding toppings:\n")

	c.Cheese = c.Ingredients.CreateCheese()
	fmt.Fprintf(outBuilder, "    %s Cheese\n", c.Cheese)

	c.InitWriter()
	if _, err := c.Writer.WriteString(outBuilder.String()); err != nil {
		return err
	}

	return nil
}
