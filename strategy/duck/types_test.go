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
		output  *bytes.Buffer
		subject duck.Duck
	)

	BeforeEach(func() {
		output = &bytes.Buffer{}
	})

	When("a MallardDuck is used", func() {

		BeforeEach(func() {
			subject = &duck.MallardDuck{
				GenericDuck: duck.GenericDuck{
					Writer: output,
				},
			}
		})

		It("can fly", func() {
			Expect(subject.PerformFly()).To(Succeed())
			Expect(output.String()).To(Equal("I'm flying!!\n"))
		})

		It("can quack", func() {
			Expect(subject.PerformQuack()).To(Succeed())
			Expect(output.String()).To(Equal("Quack\n"))
		})

		It("can swim", func() {
			Expect(subject.Swim()).To(Succeed())
			Expect(output.String()).To(Equal("All ducks float, even decoys!\n"))
		})

		It("can be displayed", func() {
			Expect(subject.Display()).To(Succeed())
			Expect(output.String()).To(Equal("I'm a real Mallard duck\n"))
		})

	})

	When("a ModelDuck is used", func() {

		BeforeEach(func() {
			subject = &duck.ModelDuck{
				GenericDuck: duck.GenericDuck{
					Writer: output,
				},
			}
		})

		It("can't fly by default", func() {
			Expect(subject.PerformFly()).To(Succeed())
			Expect(output.String()).To(Equal("I can't fly\n"))
		})

		It("can't quack by default", func() {
			Expect(subject.PerformQuack()).To(Succeed())
			Expect(output.String()).To(Equal("<< Silence >>\n"))
		})

		It("can change its flying behavior", func() {
			Expect(subject.PerformFly()).To(Succeed())
			subject.SetFlyBehavior(&behavior.FlyRocketPowered{})
			Expect(subject.PerformFly()).To(Succeed())
			outString := output.String()
			Expect(outString).To(ContainSubstring("I can't fly\n"))
			Expect(outString).To(ContainSubstring("I'm flying with a rocket!\n"))
		})

		It("can change its quacking behavior", func() {
			Expect(subject.PerformQuack()).To(Succeed())
			subject.SetQuackBehavior(&behavior.Squeak{})
			Expect(subject.PerformQuack()).To(Succeed())
			outString := output.String()
			Expect(outString).To(ContainSubstring("<< Silence >>\n"))
			Expect(outString).To(ContainSubstring("Squeak\n"))
		})

		It("can swim", func() {
			Expect(subject.Swim()).To(Succeed())
			Expect(output.String()).To(Equal("All ducks float, even decoys!\n"))
		})

		It("can be displayed", func() {
			Expect(subject.Display()).To(Succeed())
			Expect(output.String()).To(Equal("I'm a model duck\n"))
		})

	})

	When("a GenericDuck is used", func() {

		BeforeEach(func() {
			subject = &duck.GenericDuck{
				Writer: output,
			}
		})

		It("can fly", func() {
			Expect(subject.PerformFly()).To(Succeed())
			Expect(output.String()).To(Equal("I'm flying!!\n"))
		})

		It("can quack", func() {
			Expect(subject.PerformQuack()).To(Succeed())
			Expect(output.String()).To(Equal("Quack\n"))
		})

		It("can swim", func() {
			Expect(subject.Swim()).To(Succeed())
			Expect(output.String()).To(Equal("All ducks float, even decoys!\n"))
		})

		It("can be displayed", func() {
			Expect(subject.Display()).To(Succeed())
			Expect(output.String()).To(Equal("I'm a non-descript duck\n"))
		})

	})

})
