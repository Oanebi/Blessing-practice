package processor

import (
	"strconv"
	"strings"
)

func Upper(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(up)" && i > 0 {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		}

		if word[i] == "(low)" && i > 0 {
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if word[i] == "(cap)" && i > 0 {
			word[i-1] = strings.Title(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if strings.Contains(word[i], "(") && strings.Contains(word[i], ")") {
			word[i] = strings.Trim(word[i], "(),.?!:;")
			if strings.Contains(word[i], ",") {
				result := strings.Split(word[i], ",")

				num, err := strconv.Atoi(result[1])
				if err == nil {
					start := i - num
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						if result[0] == "up" {
							word[j] = strings.ToUpper(word[j])
						} else if result[0] == "low" {
							word[j] = strings.ToLower(word[j])
						} else if result[0] == "cap" {
							word[j] = strings.Title(word[j])
						}
					}
				}
				word = append(word[:i], word[i+1:]...)
				i--

			}
		}
	}
	return strings.Join(word, " ")
}
