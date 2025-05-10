package pizza_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adambkaplan/go-design-patterns/factory/pizza"
)

var _ = Describe("Basic Pizza", func() {

	PIt("bakes at 350 for 25 minutes")

	PIt("is boxed")

	PIt("is cut in diagonal slices")

	PIt("is prepared with thin crust, shredded mozarella, and marinara sauce")

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
