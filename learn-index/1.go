// Print the first and last character of the string "hello world" using indexing.
package main

import "fmt"

func index(s string) string {
	return s[len(s)-1:] + s[:1]

}

func main() {
	fmt.Println(index("hello world"))
}

// Another way

// func main() {
// 	s := "hello world"
// 	fmt.Println((string(s[0])))
// 	fmt.Println(string(s[len(s)-1]))
// }
