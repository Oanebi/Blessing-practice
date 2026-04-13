package processor

import (
	"strconv"
	"strings"
)

func Low(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(low)" && i > 0 {
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)

		}

		if strings.HasPrefix(word[i], "(") && strings.HasSuffix(word[i], ")") {
			word[i] = strings.Trim(word[i], "()")
			if strings.Contains(word[i], ",") {
				result := strings.Split(word[i], ",")
				num, err := strconv.Atoi(result[1])
				if err == nil {
					start := i - num
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						word[j] = strings.ToLower(word[j])
					}
				}

			}
			word = append(word[:i], word[i+1:]...)
		}

	}
	return strings.Join(word, " ")

}
