// Reverse a string using a loop (without using [::-1]).
package main

import "fmt"

func reverse(s string) {

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(string(s[i]))
	}
}
func main() {
	reverse("hello world")
}
