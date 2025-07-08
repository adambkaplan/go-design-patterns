package chicago_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChicago(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chicago Suite")
}
