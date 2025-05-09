package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory Pattern", func() {

	When("the pizza stores receive orders", func() {

		It("prints the process of makin the pizzas", func(ctx SpecContext) {
			output, err := RunTestHarness(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ContainSubstring("Preparing NY Style Sauce and Cheese Pizza"))
			Expect(output).To(ContainSubstring("Tossing dough..."))
			Expect(output).To(ContainSubstring("Adding sauce..."))
			Expect(output).To(ContainSubstring("Adding toppings\n    Grated Regiano Cheese"))
			Expect(output).To(ContainSubstring("Bake for 25 minuts at 350"))
			Expect(output).To(ContainSubstring("Cutting the pizza into diagonal slices"))
			Expect(output).To(ContainSubstring("Placing pizza in official PizzaStore box"))
			Expect(output).To(ContainSubstring("Ethan ordered a NY Style Sauce and Cheese Pizza"))

			Expect(output).To(ContainSubstring("Preparing Chicago Style Deep Dish Cheese Pizza"))
			Expect(output).To(ContainSubstring("Adding toppings\n    Shredded Mozarella Cheese"))
			Expect(output).To(ContainSubstring("Cutting the pizza into square slices"))
			Expect(output).To(ContainSubstring("Ethan ordered a Chicago Style Deep Dish Cheese Pizza"))
		})
	})
})
