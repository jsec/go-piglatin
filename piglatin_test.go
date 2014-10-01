package piglatin

import "testing"

func TestStartsWithConsonant(t *testing.T) {
	word := "derp"
	expected := "erpday"
	actual := Translate(word)

	if expected != actual {
		t.Error("Expected " + expected + ", got " + actual)
	}
}

func TestStartsWithVowel(t *testing.T) {
	word := "eggnog"
	expected := "eggnogway"
	actual := Translate(word)

	if expected != actual {
		t.Error("Expected " + expected + ", got " + actual)
	}
}

func TestStartsWithConsonantChunk(t *testing.T) {
	word := "chloroform"
	expected := "oroformchlay"
	actual := Translate(word)

	if expected != actual {
		t.Error("Expected " + expected + ", got " + actual)
	}
}
