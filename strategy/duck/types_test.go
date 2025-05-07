package duck_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
	"github.com/adambkaplan/go-design-patterns/strategy/duck"
)

var _ = Describe("Ducks", func() {

	var (
		output *bytes.Buffer
	)

	BeforeEach(func() {
		output = &bytes.Buffer{}
	})

	When("a MallardDuck is used", func() {

		var (
			subject *duck.MallardDuck
		)

		BeforeEach(func() {
			subject = &duck.MallardDuck{
				Writer: output,
			}
		})

		It("can fly", func() {
			subject.PerformFly()
			Expect(output.String()).To(Equal("I can fly!\n"))
		})

		It("can quack", func() {
			subject.PerformQuack()
			Expect(output.String()).To(Equal("Quack!\n"))
		})

		It("can swim", func() {
			subject.Swim()
			Expect(output.String()).To(Equal("All ducks can swim, even decoys!\n"))
		})

		It("can be displayed", func() {
			subject.Display()
			Expect(output.String()).To(Equal("I'm a Mallard duck.\n"))
		})

	})

	When("a ModelDuck is used", func() {

		var (
			subject *duck.ModelDuck
		)

		BeforeEach(func() {
			subject = &duck.ModelDuck{
				Writer: output,
			}
		})

		It("can't fly by default", func() {
			subject.PerformFly()
			Expect(output.String()).To(Equal("I can't fly\n"))
		})

		It("can't quack by default", func() {
			subject.PerformQuack()
			Expect(output.String()).To(Equal("<< Silent >>\n"))
		})

		It("can change its flying behavior", func() {
			subject.PerformFly()
			subject.FlyBehavior = &behavior.FlyRocketPowered{}
			subject.PerformFly()
			outString := output.String()
			Expect(outString).To(ContainSubstring("I can't fly"))
			Expect(outString).To(ContainSubstring("I'm flying with a rocket!\n"))
		})

		It("can change its quacking behavior", func() {
			subject.PerformQuack()
			subject.QuackBehavior = &behavior.Squeak{}
			subject.PerformQuack()
			outString := output.String()
			Expect(outString).To(ContainSubstring("I can't fly"))
			Expect(outString).To(ContainSubstring("I'm flying with a rocket!\n"))
		})

		It("can swim", func() {
			subject.Swim()
			Expect(output.String()).To(Equal("All ducks can swim, even decoys!\n"))
		})

		It("can be displayed", func() {
			subject.Display()
			Expect(output.String()).To(Equal("I'm a model duck.\n"))
		})

	})
})
