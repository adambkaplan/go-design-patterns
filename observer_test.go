package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Observer Pattern", func() {

	When("the weather station is run", func() {

		It("prints the weather displays to stdout", func(ctx SpecContext) {
			output, err := RunTestHarness(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ContainSubstring("Current conditions: "))
			Expect(output).To(ContainSubstring("Forecast: Improving weather on the way!"))
			Expect(output).To(ContainSubstring("Forecast: Watch out for rainy weather"))
			Expect(output).To(ContainSubstring("Forecast: More of the same"))
		})
	})
})
