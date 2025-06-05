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
	outBuilder.WriteString(fmt.Sprintf("Preparing %s\n", c.Name))

	c.Dough = c.Ingredients.CreateDough()
	outBuilder.WriteString(fmt.Sprintf("Tossing %s...\n", c.Dough))

	c.Sauce = c.Ingredients.CreateSauce()
	outBuilder.WriteString(fmt.Sprintf("Adding %s...\n", c.Sauce))

	outBuilder.WriteString("Adding toppings:\n")

	c.Cheese = c.Ingredients.CreateCheese()
	outBuilder.WriteString(fmt.Sprintf("    %s Cheese\n", c.Cheese))

	c.InitWriter()
	if _, err := c.Writer.WriteString(outBuilder.String()); err != nil {
		return err
	}

	return nil
}
