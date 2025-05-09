package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decorator pattern", func() {

	When("the coffee shop example is run", func() {

		It("prints the coffee prices", func(ctx SpecContext) {
			output, err := RunTestHarness(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ContainSubstring("Espresso $1.99"))
			Expect(output).To(ContainSubstring("Dark Roast Coffee, Mocha, Mocha, Whip $1.49"))
			Expect(output).To(ContainSubstring("House Blend Coffee, Soy, Mocha, Whip $1.34"))
		})
	})
})
