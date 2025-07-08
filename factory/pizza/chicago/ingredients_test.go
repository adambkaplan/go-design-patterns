package chicago_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adambkaplan/go-design-patterns/factory/pizza"
	"github.com/adambkaplan/go-design-patterns/factory/pizza/chicago"
)

var _ = Describe("Chicago Ingredients", func() {

	var ingredients pizza.IngredientFactory

	BeforeEach(func() {
		ingredients = chicago.ChicagoIngredients{}
	})

	It("creates Plum Tomato sauce", func() {
		Expect(ingredients.CreateSauce()).To(Equal(pizza.PlumTomato))
	})

	It("creates Shredded Mozarella cheese", func() {
		Expect(ingredients.CreateCheese()).To(Equal(pizza.ShreddedMozarella))
	})
})
