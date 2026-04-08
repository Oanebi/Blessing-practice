//Count how many times the letter 'l' appears in "hello world" using a loop.

package main

import (
	"fmt"
	"strings"
)

//	func countletter(s string, r rune) int {
//		count := 0
//		for _, ch := range s {
//			if ch == r {
//				count++
//			}
//		}
//		return count
//	}
func main() {
	sentence := "hello world"
	s := strings.Count(sentence, "o")
	fmt.Println(s)
}
