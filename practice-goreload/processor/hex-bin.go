package processor

import (
	"strconv"
	"strings"
)

func Hex(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" && i > 0 {
			num, err := strconv.ParseInt(word[i-1], 16, 64)
			if err == nil {
				word[i-1] = strconv.FormatInt(int64(num), 10)
				word = append(word[:i], word[i+1:]...)

			}
		}
	}
	return strings.Join(word, " ")

}

func Bin(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(bin)" && i > 0 {
			num, err := strconv.ParseInt(word[i-1], 2, 64)
			if err == nil {
				word[i-1] = strconv.FormatInt(int64(num), 10)
				word = append(word[:i], word[i+1:]...)

			}
		}
	}
	return strings.Join(word, " ")

}
