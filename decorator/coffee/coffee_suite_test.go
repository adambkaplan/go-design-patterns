package coffee_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCoffee(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Coffee Suite")
}
