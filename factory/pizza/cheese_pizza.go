package pizza

type CheesePizza struct {
	BasicPizza

	Ingredients IngredientFactory
}

func (c *CheesePizza) Prepare() error {
	c.BasicPizza.InitWriter()
	if _, err := c.BasicPizza.Writer.WriteString("Preparing " + c.GetName() + "\n"); err != nil {
		return err
	}

	c.BasicPizza.Dough = c.Ingredients.CreateDough()
	c.BasicPizza.Sauce = c.Ingredients.CreateSauce()
	c.BasicPizza.Cheese = c.Ingredients.CreateCheese()
	return nil
}
