package main_test

import (
	"context"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoDesignPatterns(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoDesignPatterns Suite")
}

func RunTestHarness(ctx context.Context) (string, error) {
	runCmd := exec.CommandContext(ctx, "./go-design-patterns")
	output, err := runCmd.CombinedOutput()
	return string(output), err
}

var _ = BeforeSuite(func(ctx SpecContext) {
	makeCmd := exec.CommandContext(ctx, "make", "build")
	Expect(makeCmd.Run()).To(Succeed())
})
