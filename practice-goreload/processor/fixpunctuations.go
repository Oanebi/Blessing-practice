package processor

import (
	"regexp"
)

func Fixpunctuations(s string) string {
	pattern := regexp.MustCompile(`\s+([,!?;.]+)`)
	s = pattern.ReplaceAllString(s, "$1 ")

	return s
}
