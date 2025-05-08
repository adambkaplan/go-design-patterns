package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strategy Pattern", func() {

	When("the duck simulator is run", func() {

		It("prints the behavior to stdout", func(ctx SpecContext) {
			output, err := RunTestHarness(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ContainSubstring("Quack\n"))
			Expect(output).To(ContainSubstring("I'm flying!!\n"))
			Expect(output).To(ContainSubstring("I can't fly\n"))
			Expect(output).To(ContainSubstring("I'm flying with a rocket!\n"))
		})

	})
})
