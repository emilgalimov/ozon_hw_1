package romannumerals

import "strings"

type RomanNumber struct {
	a int
	w string
}

var dict = []RomanNumber{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func Encode(n int) (string, bool) {
	if n < 1 {
		return "", false
	}

	var result string

	for {
		if n == 0 {
			break
		}
		for _, v := range dict {
			if v.a > n {
				continue
			}
			n = n - v.a
			result += v.w
			break
		}
	}

	return result, true
}

func Decode(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	var result int

	for {
		if s == "" {
			break
		}

		var hasWord bool
		for _, v := range dict {
			if !strings.HasPrefix(s, v.w) {
				continue
			}
			s = strings.TrimPrefix(s, v.w)
			result += v.a
			hasWord = true
			break
		}

		if !hasWord {
			return 0, false
		}
	}

	return result, true
}
