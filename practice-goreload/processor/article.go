package processor

import (
	"strings"
)

func Article(s string) string {
	word := strings.Fields(s)

	vowel := "aeiouAEIOU"
	for i := 0; i < len(word)-1; i++ {
		if word[i] == "a" && strings.Contains(vowel, word[i+1][:1]) {
			word[i] = "an"
		} else if word[i] == "A" && strings.Contains(vowel, word[i+1][:1]) {
			word[i] = "An"

		} else if word[i] == "an" && !strings.Contains(vowel, word[i+1][:1]) {
			word[i] = "a"
		} else if word[i] == "An" && !strings.Contains(vowel, word[i+1][:1]) {
			word[i] = "A"
		} else if word[i+1] == "hour" {
			word[i] = "an"
		}
	}
	return strings.Join(word, " ")
}
