package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hello World", func() {

	When("the test is run", func() {

		It("prints Hello World", func(ctx SpecContext) {

			output, err := RunTestHarness(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ContainSubstring("Hello, Head First Design Patterns!\n"))
		})
	})
})
