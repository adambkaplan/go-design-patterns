package ny_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NY Suite")
}
