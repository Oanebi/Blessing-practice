// Given s = "abcdef", print the character at index 2 and the character at the second-to-last position.
package main

import "fmt"

func main() {
	s := "abcdef"
	fmt.Println(string(s[2]))
	fmt.Println(string(s[len(s)-2]))

}
