package piglatin

import (
	"strings"
)

const (
	vowel            string = "aeiou"
	baseSuffix       string = "ay"
	vowelStartSuffix string = "w" + baseSuffix
)

func Translate(word string) string {

	if endsInPunctuation(word) {
		punc := word[len(word)-1:]
		stripped := word[:len(word)-1]

		return getTranslation(stripped) + punc
	} else {
		return getTranslation(word)
	}
}

func getTranslation(word string) string {
	first := word[0:1]

	if strings.Contains(vowel, first) {
		return word + vowelStartSuffix
	} else {
		firstVowel := findFirstVowel(word)
		return word[firstVowel:] + word[:firstVowel] + baseSuffix
	}
}

func findFirstVowel(word string) int {
	for index, letter := range word {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u':
			return index
		}
	}

	return 0
}

func endsInPunctuation(word string) bool {
	char := word[len(word)-1:]
	switch char {
	case ".", ",", "!", "?", ":", ";":
		return true
	}

	return false
}
