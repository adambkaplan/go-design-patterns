package duck_test

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adambkaplan/go-design-patterns/strategy/behavior"
	"github.com/adambkaplan/go-design-patterns/strategy/duck"
)

var _ = Describe("Ducks", func() {

	var (
		origStdOut    *os.File
		stdOutCapture *os.File
		err           error
	)

	BeforeEach(func() {
		origStdOut = os.Stdout
		stdOutCapture, err = os.CreateTemp("", "test-duck-*.log")
		Expect(err).NotTo(HaveOccurred())
		os.Stdout = stdOutCapture
	})

	AfterEach(func() {
		os.Stdout = origStdOut
		Expect(os.Remove(stdOutCapture.Name())).To(Succeed())
	})

	When("a Mallard Duck is used", func() {

		var subject *duck.MallardDuck

		BeforeEach(func() {
			subject = duck.NewMallardDuck()
		})

		It("can fly with wings", func() {
			subject.PerformFly()
			output, err := os.ReadFile(stdOutCapture.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeEmpty())
			Expect(string(output)).To(Equal("I'm flying with wings!\n"))
		})

		It("can quack", func() {
			subject.PerformQuack()
			output, err := os.ReadFile(stdOutCapture.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeEmpty())
			Expect(string(output)).To(Equal("Quack!\n"))
		})

		It("can swim", func() {
			subject.Swim()
			output, err := os.ReadFile(stdOutCapture.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeEmpty())
			Expect(string(output)).To(Equal("All ducks float, even decoys!\n"))
		})

		It("can be displayed", func() {
			subject.Display()
			output, err := os.ReadFile(stdOutCapture.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeEmpty())
			Expect(string(output)).To(Equal("I'm a real Mallard duck\n"))
		})
	})

	When("a Model Duck is used", func() {

		var subject *duck.ModelDuck

		BeforeEach(func() {
			subject = duck.NewModelDuck()
		})

		It("can't fly by default", func() {
			subject.PerformFly()
			output, err := os.ReadFile(stdOutCapture.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeEmpty())
			Expect(string(output)).To(Equal("I can't fly\n"))
		})

		It("has a silent quack", func() {
			subject.PerformQuack()
			output, err := os.ReadFile(stdOutCapture.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeEmpty())
			Expect(string(output)).To(Equal("<< Silence >>\n"))
		})

		It("can change its flying behavior", func() {
			subject.PerformFly()
			subject.FlyBehavior = behavior.FlyRocketPowered{}
			subject.PerformFly()

			scanner := bufio.NewScanner(stdOutCapture)

			prevLine := ""
			for scanner.Scan() {
				line := scanner.Text()
				Expect(line).NotTo(Equal(prevLine))
				prevLine = line
			}
			Expect(scanner.Err()).NotTo(HaveOccurred())
		})

		It("can change its quacking behavior", func() {
			subject.PerformQuack()
			subject.QuackBehavior = behavior.Squeak{}
			subject.PerformQuack()
			scanner := bufio.NewScanner(stdOutCapture)

			prevLine := ""
			for scanner.Scan() {
				line := scanner.Text()
				Expect(line).NotTo(Equal(prevLine))
				prevLine = line
			}
			Expect(scanner.Err()).NotTo(HaveOccurred())
		})

		It("can swim", func() {
			subject.Swim()
			output, err := os.ReadFile(stdOutCapture.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeEmpty())
			Expect(string(output)).To(Equal("All ducks float, even decoys!\n"))
		})

		It("can be displayed", func() {
			subject.Display()
			output, err := os.ReadFile(stdOutCapture.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(output).NotTo(BeEmpty())
			Expect(string(output)).To(Equal("I'm a model duck\n"))
		})
	})
})
