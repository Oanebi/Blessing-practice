//Write a function that checks if two strings are anagrams of each other using only loops.

package main

import "fmt"

func anagrams(s1, s2 string) string {
	if len(s1) != len(s2) {
		return "not anagram"
	}
	count := [26]int{}
	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']--
		count[s2[i]-'a']++

	}
	for _, v := range count {
		if v != 0 {
			return " not anagram"
		}
	}
	return " anagram"
}
func main() {
	fmt.Println(anagrams("silent", "listen"))
	fmt.Println(anagrams("hello", "world"))

}
