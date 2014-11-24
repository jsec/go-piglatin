package piglatin_test

import (
	. "github.com/jsec/piglatin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Translations", func() {
	It("starts with a consonant", func() {
		word := "derp"
		Expect(Translate(word)).To(Equal("erpday"))
	})

	It("starts with a vowel", func() {
		word := "eggnog"
		Expect(Translate(word)).To(Equal("eggnogway"))
	})

	It("starts with a consonant chunk", func() {
		word := "chloroform"
		Expect(Translate(word)).To(Equal("oroformchlay"))
	})
})
