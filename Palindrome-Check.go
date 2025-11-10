package main

import (
	"fmt"
	"strings"
	"unicode"
)

func PalindromeCheck(s string) bool {
	lowerText := strings.ToLower(s)
	var cleaned strings.Builder
	var reversed strings.Builder

	for _, char := range lowerText {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(char)
		}
	}

	cleanedtext := []rune(cleaned.String())

	for i := len(cleanedtext) - 1; i >= 0; i-- {
		reversed.WriteRune(cleanedtext[i])
	}

	if reversed.String() == string(cleanedtext) {
		return true
	}
	return false

}
func main() {
	text1 := "sdfghjjhgfgd"
	text2 := "m om"

	fmt.Println(PalindromeCheck(text1)) // false
	fmt.Println(PalindromeCheck(text2)) // true
}
