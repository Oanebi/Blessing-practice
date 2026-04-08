// Given a sentence, count the number of words that start with a capital letter using a loop.

package main

import (
	"fmt"
	"strings"
)

func count(s string) int {
	count := 0
	word := strings.Fields(s)
	for _, word := range word {
		if word >= "A" || word <= "Z" {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(count(" Hello World Doing Today yes"))
}
