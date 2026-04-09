// Given a string, find the first non-repeating character using a loop. Return its index.
package main

import "fmt"

func nonRepeating(s string) int {
	for i := 0; i < len(s); i++ {
		count := 0
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				count++
			}
		}
		if count == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(nonRepeating("helloworld"))
	fmt.Println(nonRepeating("aabb"))
}
