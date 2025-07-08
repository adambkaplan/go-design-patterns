package pizza_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adambkaplan/go-design-patterns/factory/pizza"
	"github.com/adambkaplan/go-design-patterns/factory/pizza/ny"
)

var _ = Describe("Cheese Pizza", func() {

	var (
		cheesePizza *pizza.CheesePizza
		output      *bytes.Buffer
	)

	BeforeEach(func() {
		output = &bytes.Buffer{}
		cheesePizza = &pizza.CheesePizza{
			BasicPizza: pizza.BasicPizza{
				Name:   "New York Cheese Pizza",
				Writer: output,
			},
		}
	})

	It("bakes at 350 for 25 minutes", func() {
		Expect(cheesePizza.Bake()).To(Succeed())
		Expect(output.String()).To(Equal("Bake for 25 minutes at 350\n"))
	})

	It("is boxed", func() {
		Expect(cheesePizza.Box()).To(Succeed())
		Expect(output.String()).To(Equal("Placing pizza in official PizzaStore box\n"))
	})

	It("is cut in diagonal slices", func() {
		Expect(cheesePizza.Cut()).To(Succeed())
		Expect(output.String()).To(Equal("Cutting the pizza into diagonal slices\n"))
	})

	It("fails to prepare if no ingredients are specified", func() {
		Expect(cheesePizza.Prepare()).To(Not(Succeed()))
	})

	It("prepares with NY ingredients", func() {
		cheesePizza.Ingredients = ny.NYIngredients{}
		Expect(cheesePizza.Prepare()).To(Succeed())
		actualOutput := output.String()
		Expect(actualOutput).To(ContainSubstring("Preparing New York Cheese Pizza\n"))
		Expect(actualOutput).To(ContainSubstring("Tossing %s", pizza.ThinCrust))
		Expect(actualOutput).To(ContainSubstring("Adding %s", pizza.Marinara))
		Expect(actualOutput).To(ContainSubstring("    %s", pizza.ShreddedMozarella))
	})

})
