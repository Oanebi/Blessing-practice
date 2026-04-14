package processor

import (
	"regexp"
)

func Quote(s string) string {
	pattern := regexp.MustCompile(`^'\s+`)
	s = pattern.ReplaceAllString(s, "'")
	pattern1 := regexp.MustCompile(`\s+'\s+`)
	s = pattern1.ReplaceAllString(s, " '")
	pattern2 := regexp.MustCompile(`\s+([.?!])\s+'`)
	s = pattern2.ReplaceAllString(s, "$1' ")
	return s

}
