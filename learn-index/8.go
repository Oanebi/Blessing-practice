// Write a loop that replaces every vowel in a string with '*'.
package main

import (
	"fmt"
	"strings"
)

func replacevowel(s string) string {
	result := ""
	vowels := "aeiou"
	for _, v := range s {
		if strings.ContainsRune(vowels, v) {
			result += "*"
		} else {
			result += string(v)
		}

	}
	return result
}
func main() {
	fmt.Println(replacevowel("houses"))
}
