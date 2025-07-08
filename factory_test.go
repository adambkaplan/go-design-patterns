package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory Pattern", func() {

	When("the pizza stores receive orders", func() {

		It("prints the process of making the pizzas", func(ctx SpecContext) {
			output, err := RunTestHarness(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ContainSubstring("Preparing New York Style Cheese Pizza\n"))
			Expect(output).To(ContainSubstring("Tossing Thin Crust Dough...\n"))
			Expect(output).To(ContainSubstring("Adding Marinara Sauce...\n"))
			Expect(output).To(ContainSubstring("Adding toppings:\n    Shredded Mozarella Cheese\n"))
			Expect(output).To(ContainSubstring("Bake for 25 minutes at 350\n"))
			Expect(output).To(ContainSubstring("Cutting the pizza into diagonal slices\n"))
			Expect(output).To(ContainSubstring("Placing pizza in official PizzaStore box\n"))
			Expect(output).To(ContainSubstring("Ethan ordered a New York Style Cheese Pizza\n"))

			Expect(output).To(ContainSubstring("Preparing Chicago Style Deep Dish Cheese Pizza\n"))
			// Expect(output).To(ContainSubstring("Adding toppings\n    Shredded Mozarella Cheese"))
			// Expect(output).To(ContainSubstring("Cutting the pizza into square slices"))
			Expect(output).To(ContainSubstring("Joel ordered a Chicago Style Deep Dish Cheese Pizza\n"))
		})
	})
})
