package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPgfib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pgfib Suite")
}
