package coffee_test

import (
	"github.com/adambkaplan/go-design-patterns/decorator/coffee"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Condiments", func() {

	var testBeverage coffee.Beverage

	BeforeEach(func() {
		testBeverage = &coffee.HouseBlend{}
	})

	When("Milk is added", func() {
		var condiment *coffee.Milk

		BeforeEach(func() {
			condiment = &coffee.Milk{}
		})

		It("should add the correct price to a beverage", func() {
			condiment.Beverage = testBeverage
			expected := testBeverage.Cost() + 0.10
			Expect(condiment.Cost()).To(Equal(expected))
		})

		It("should add the correct description to a beverage", func() {
			condiment.Beverage = testBeverage
			expected := testBeverage.Description() + ", Milk"
			Expect(condiment.Description()).To(Equal(expected))
		})

		It("should work if no beverage is set", func() {
			Expect(condiment.Cost()).To(Equal(0.10))
			Expect(condiment.Description()).To(Equal("unknown, Milk"))
		})
	})

	When("Mocha is added", func() {
		var condiment *coffee.Mocha

		BeforeEach(func() {
			condiment = &coffee.Mocha{}
		})

		It("should add the correct price to a beverage", func() {
			condiment.Beverage = testBeverage
			expected := testBeverage.Cost() + 0.20
			Expect(condiment.Cost()).To(Equal(expected))
		})

		It("should add the correct description to a beverage", func() {
			condiment.Beverage = testBeverage
			expected := testBeverage.Description() + ", Mocha"
			Expect(condiment.Description()).To(Equal(expected))
		})

		It("should work if no beverage is set", func() {
			Expect(condiment.Cost()).To(Equal(0.20))
			Expect(condiment.Description()).To(Equal("unknown, Mocha"))
		})
	})

	When("Soy is added", func() {
		var condiment *coffee.Soy

		BeforeEach(func() {
			condiment = &coffee.Soy{}
		})

		It("should add the correct price to a beverage", func() {
			condiment.Beverage = testBeverage
			expected := testBeverage.Cost() + 0.15
			Expect(condiment.Cost()).To(Equal(expected))
		})

		It("should add the correct description to a beverage", func() {
			condiment.Beverage = testBeverage
			expected := testBeverage.Description() + ", Soy"
			Expect(condiment.Description()).To(Equal(expected))
		})

		It("should work if no beverage is set", func() {
			Expect(condiment.Cost()).To(Equal(0.15))
			Expect(condiment.Description()).To(Equal("unknown, Soy"))
		})
	})

	When("Whip is added", func() {

		var condiment *coffee.Whip

		BeforeEach(func() {
			condiment = &coffee.Whip{}
		})

		It("should add the correct price to a beverage", func() {
			condiment.Beverage = testBeverage
			expected := testBeverage.Cost() + 0.10
			Expect(condiment.Cost()).To(Equal(expected))
		})

		It("should add the correct description to a beverage", func() {
			condiment.Beverage = testBeverage
			expected := testBeverage.Description() + ", Whip"
			Expect(condiment.Description()).To(Equal(expected))
		})

		It("should work if no beverage is set", func() {
			Expect(condiment.Cost()).To(Equal(0.10))
			Expect(condiment.Description()).To(Equal("unknown, Whip"))
		})
	})
})
