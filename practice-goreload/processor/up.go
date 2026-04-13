package processor

import (
	"strconv"
	"strings"
)

func Upper(s string) string {
	text := strings.Fields(s)
	for i := 0; i < len(text); i++ {
		if text[i] == "(up)" && i > 0 {
			text[i-1] = strings.ToUpper(text[i-1])

			text = append(text[:i], text[i+1:]...)
			i--
		}
		if strings.HasPrefix(text[i], "(") && strings.HasSuffix(text[i], ")") { // (up,2)
			text[i] = strings.Trim(text[i], "()")
			if strings.Contains(text[i], ",") {
				result := strings.Split(text[i], ",")
				num, err := strconv.Atoi(result[1])
				if err == nil {
					start := i - num
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						text[j] = strings.ToUpper(text[j])
					}
				}
			}
			text = append(text[:i], text[i+1:]...)
			i++
		}

	}
	return strings.Join(text, " ")
}
