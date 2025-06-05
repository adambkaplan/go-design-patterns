package pizza_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adambkaplan/go-design-patterns/factory/pizza"
)

var _ = Describe("Basic Pizza", func() {

	var (
		basicPizza *pizza.BasicPizza
		output     *bytes.Buffer
	)

	BeforeEach(func() {
		output = &bytes.Buffer{}
		basicPizza = &pizza.BasicPizza{
			Writer: output,
		}
	})

	It("bakes at 350 for 25 minutes", func() {
		Expect(basicPizza.Bake()).To(Succeed())
		Expect(output.String()).To(Equal("Bake for 25 minutes at 350\n"))
	})

	It("is boxed", func() {
		Expect(basicPizza.Box()).To(Succeed())
		Expect(output.String()).To(Equal("Placing pizza in official PizzaStore box\n"))
	})

	It("is cut in diagonal slices", func() {
		Expect(basicPizza.Cut()).To(Succeed())
		Expect(output.String()).To(Equal("Cutting the pizza into diagonal slices\n"))
	})

	It("does nothing during preparation", func() {
		Expect(basicPizza.Prepare()).To(Succeed())
		Expect(output.String()).To(BeEmpty())
	})

})

var _ = Describe("Basic Store", func() {

	It("creates a Basic Pizza by default", func() {
		output := &bytes.Buffer{}
		store := &pizza.BasicStore{
			Writer: output,
		}
		pizza, err := store.OrderPizza("cheese")
		Expect(err).NotTo(HaveOccurred())
		Expect(pizza.GetName()).To(Equal("Basic Pizza"))
	})
})
