package reverseparentheses

import (
	"regexp"
	"strings"
)

const _pattern = `[\(]([^\(|^\)]*)[\)]`

func Reverse(s string) string {
	return reversion(s)
}

func reversion(s string) string {
	if !hasWrap(s) {
		return s
	}
	return reversion(removeInnerWraps(s))
}

func hasWrap(s string) bool {
	return strings.Contains(s, "(")
}

func removeInnerWraps(s string) string {
	r, _ := regexp.Compile(_pattern)

	for _, match := range r.FindStringSubmatch(s) {
		s = strings.Replace(s, match, reverse(match), 1)
	}

	return s
}

func reverse(s string) string {
	s = strings.Replace(s, "(", "", 1)
	s = strings.Replace(s, ")", "", 1)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
