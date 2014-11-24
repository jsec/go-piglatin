package piglatin_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPiglatin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Piglatin Suite")
}
