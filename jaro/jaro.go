package jaro

import (
	"math"
	"strings"
)

func Distance(word1 string, word2 string) float64 {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	s1 := len(word1)
	s2 := len(word2)

	maxBetween := math.Floor(math.Max(float64(s1), float64(s2))/2) - 1
	var m int

	if word1 == word2 {
		return 1
	}

	str1Matches := make([]bool, len(word1))
	str2Matches := make([]bool, len(word2))
	for i, letter := range word1 {
		isM := false
		for i2, letter2 := range word2 {
			between := math.Abs(float64(i) - float64(i2))
			if letter == letter2 && between <= maxBetween {
				if str2Matches[i2] {
					continue
				}
				str1Matches[i] = true
				str2Matches[i2] = true
				isM = true
				break
			}

		}
		if isM {
			m++
		}
	}

	k := 0
	transpositions := 0
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

	tf := float64(transpositions) / float64(2)
	mf := float64(m)
	s1f := float64(s1)
	s2f := float64(s2)

	return (float64(1) / float64(3)) * ((mf / s1f) + (mf / s2f) + ((mf - tf) / mf))
}
