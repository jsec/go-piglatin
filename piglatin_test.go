package piglatin

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)

	VerifyTranslation := func(word, expected string) {
		g.Assert(Translate(word)).Equal(expected)
	}

	g.Describe("Translations", func() {
		g.Describe("First character conditions", func() {
			g.It("starts with a consonant", func() {
				VerifyTranslation("derp", "erpday")
			})

			g.It("starts with a vowel", func() {
				VerifyTranslation("eggnog", "eggnogway")
			})

			g.It("starts with a consonant chunk", func() {
				VerifyTranslation("chloroform", "oroformchlay")
			})
		})

		g.Describe("Ends in punctuation", func() {
			g.It("ends with punctuation", func() {
				VerifyTranslation("herp,", "erphay,")
			})

			g.It("ends with a period", func() {
				VerifyTranslation("derp.", "erpday.")
			})
		})
	})
}
