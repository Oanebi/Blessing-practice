// Loop through the string "python" and print each character on a new line.
package main

import "fmt"

func main() {
	s := "hello"
	for i, ch := range s {
		fmt.Println(string(ch), i)
	}
}
