package secretmessage

import (
	"sort"
)

type letter struct {
	ch    rune
	count int
}

// Decode func
func Decode(encoded string) string {

	count := make(map[rune]letter)

	//count items
	for _, char := range encoded {
		item := count[char]

		item.ch = char
		item.count++

		count[char] = item
	}

	//sort items
	var sortedCount []letter

	for _, v := range count {
		sortedCount = append(sortedCount, v)
	}

	sort.Slice(sortedCount, func(i, j int) bool {
		return sortedCount[i].count > sortedCount[j].count
	})

	//drop characters
	var resultString string

	for _, v := range sortedCount {
		if v.ch == '_' {
			break
		}
		resultString += string(v.ch)
	}

	return resultString
}
