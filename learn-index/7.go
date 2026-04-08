// Print only the characters at even indices of a given string.
package main

import (
	"fmt"
	"strings"
)

func evenindex(s string) string {
	word := " "
	for i, ch := range s {
		if i%2 == 0 {
			word += strings.ToUpper(string(ch))
		}
	}
	return word
}
func main() {
	fmt.Println(evenindex("helloo"))
}
