package piglatin

import "testing"

func TestStartsWithConsonant(t *testing.T) {
	VerifyTranslation("derp", "erpday", t)
}

func TestStartsWithVowel(t *testing.T) {
	VerifyTranslation("eggnog", "eggnogway", t)
}

func TestStartsWithConsonantChunk(t *testing.T) {
	VerifyTranslation("chloroform", "oroformchlay", t)
}

func VerifyTranslation(word, expected string, t *testing.T) {
	result := Translate(word)

	if expected != result {
		t.Error("Expected " + expected + ", got " + result)
	}
}
