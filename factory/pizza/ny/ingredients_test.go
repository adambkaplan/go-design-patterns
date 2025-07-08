package ny_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adambkaplan/go-design-patterns/factory/pizza"
	"github.com/adambkaplan/go-design-patterns/factory/pizza/ny"
)

var _ = Describe("NY Ingredients", func() {

	var ingredients pizza.IngredientFactory

	BeforeEach(func() {
		ingredients = ny.NYIngredients{}
	})

	It("creates Marinara sauce", func() {
		Expect(ingredients.CreateSauce()).To(Equal(pizza.Marinara))
	})

	It("creates Shredded Mozarella cheese", func() {
		Expect(ingredients.CreateCheese()).To(Equal(pizza.ShreddedMozarella))
	})
})
