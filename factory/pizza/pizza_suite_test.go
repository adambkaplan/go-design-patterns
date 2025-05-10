package pizza_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPizza(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pizza Suite")
}
