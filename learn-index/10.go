// Find the longest word in a sentence using a loop and indexing. No built-in max().
package main

import (
	"fmt"
	"strings"
)

func longest(s string) string {
	words := strings.Fields(s)
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func main() {
	fmt.Println(longest("today is a good day")) // "today"
}
