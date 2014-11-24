package piglatin

import (
	"fmt"
	"strings"
)

const (
	vowel            string = "aeiou"
	baseSuffix       string = "ay"
	vowelStartSuffix string = "w" + baseSuffix
)

func Translate(word string) string {
	var translatedWord string

	// how the string is translated depends on if the first letter is a
	// consonant or a vowel
	first := word[0:1]

	if strings.Contains(vowel, first) {
		translatedWord = word + vowelStartSuffix
	} else {
		firstVowel := findFirstVowel(word)
		translatedWord = word[firstVowel:] + word[:firstVowel] + baseSuffix
	}

	if endsinPunctuation(word) {
		fmt.Println(word)
	}

	return translatedWord
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

func endsinPunctuation(word string) bool {
	char := word[len(word)-1:]
	switch char {
	case ".", ",", "!", "?", ":", ";":
		return true
	}

	return false
}
