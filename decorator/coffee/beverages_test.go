package coffee_test

import (
	"github.com/adambkaplan/go-design-patterns/decorator/coffee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Beverages", func() {

	var beverage coffee.Beverage

	When("a dark roast is ordered", func() {

		BeforeEach(func() {
			beverage = &coffee.DarkRoast{}
		})

		It("is called Dark Roast Coffee", func() {
			Expect(beverage.Description()).To(Equal("Dark Roast Coffee"))
		})

		It("costs 99 cents", func() {
			Expect(beverage.Cost()).To(Equal(0.99))
		})
	})

	When("a decaf is ordered", func() {

		BeforeEach(func() {
			beverage = &coffee.Decaf{}
		})

		It("is called Decaf Coffee", func() {
			Expect(beverage.Description()).To(Equal("Decaf Coffee"))
		})

		It("costs 1.05 dollars", func() {
			Expect(beverage.Cost()).To(Equal(1.05))
		})
	})

	When("an espresso  is ordered", func() {

		BeforeEach(func() {
			beverage = &coffee.Espresso{}
		})

		It("is called Espresso", func() {
			Expect(beverage.Description()).To(Equal("Espresso"))
		})

		It("costs 1.99 dollars", func() {
			Expect(beverage.Cost()).To(Equal(1.99))
		})
	})

	When("a house blend is ordered", func() {

		BeforeEach(func() {
			beverage = &coffee.HouseBlend{}
		})

		It("is called House Blend Coffee", func() {
			Expect(beverage.Description()).To(Equal("House Blend Coffee"))
		})

		It("costs 89 cents", func() {
			Expect(beverage.Cost()).To(Equal(0.89))
		})
	})

})
