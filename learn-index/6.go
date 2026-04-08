// Check if a string is a palindrome using indexing (e.g. "racecar" should return True).

package main

import (
	"fmt"
	"strings"
)

func palindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func main() {

	fmt.Println(palindrome("racecar"))
	fmt.Println(palindrome("hello"))
	fmt.Println(palindrome("sanctus"))
	fmt.Println(palindrome("Madam"))

}
