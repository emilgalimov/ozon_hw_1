package jaro

import (
	"math"
	"strings"
)

func Distance(word1 string, word2 string) float64 {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	s1 := float64(len(word1))
	s2 := float64(len(word2))

	maxBetween := math.Floor(math.Max(s1, s2)/2) - 1

	if word1 == word2 {
		return 1
	}

	str1Matches := make([]bool, int(s1), int(s1))
	str2Matches := make([]bool, int(s2), int(s2))

	var m float64

	for i, letter := range word1 {
		var isM bool
		for i2, letter2 := range word2 {

			if letter != letter2 || math.Abs(float64(i)-float64(i2)) > maxBetween {
				continue
			}

			if str2Matches[i2] {
				continue
			}

			str1Matches[i] = true
			str2Matches[i2] = true
			isM = true
			break

		}
		if isM {
			m++
		}
	}

	var k int
	var transpositions float64

	for i := range word1 {
		if !str1Matches[i] {
			continue
		}
		for !str2Matches[k] {
			k++
		}
		if word1[i] != word2[k] {
			transpositions++
		}
		k++
	}

	if m == 0 {
		return 0
	}

	t := transpositions / 2.0

	return (1.0 / 3.0) * ((m / s1) + (m / s2) + ((m - t) / m))
}
