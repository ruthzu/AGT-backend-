package main

import (
	"fmt"
	"strings"
	"unicode"
)

func frerquencyCount(s string) map[string]int {
	lowerString := strings.ToLower(s)
	var cleanedString strings.Builder
	counter := make(map[string]int)
	for _, char := range lowerString {
		if unicode.IsLetter(char) {
			cleanedString.WriteRune(char)
		}
		words := strings.Fields(cleanedString.String())
		for _, word := range words {

			counter[word]++
		}
	}
	return counter
}

func main() {
	s := "asdfg asfdg sf"
	result := frerquencyCount(s)
	fmt.Println(result)
	fmt.Println("sdfg")
}
