package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jsec/go-piglatin"
)

func main() {
	var translatedWords []string

	// don't include the command name in word list
	wordList := os.Args[1:len(os.Args)]

	for _, word := range wordList {
		translatedWords = append(translatedWords, piglatin.Translate(word))
	}

	fmt.Println(strings.Join(translatedWords, " "))
}
