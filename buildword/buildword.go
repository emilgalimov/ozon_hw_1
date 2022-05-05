package buildword

import (
	"errors"
	"strings"
)

func BuildWord(word string, fragments []string) int {

	variants, err := getPossibleVariants(word, fragments)

	if err != nil {
		return 0
	}

	min := 0
	for _, variant := range variants {
		if min == 0 || min > len(variant) {
			min = len(variant)
		}
	}

	return min
}

func getPossibleVariants(word string, fragments []string) ([][]string, error) {

	var returnPossibilities [][]string
	for _, fragment := range fragments {
		if 0 != strings.Index(word, fragment) {
			continue
		}
		afterTrim := strings.TrimPrefix(word, fragment)
		if "" == afterTrim {
			returnPossibilities = append(returnPossibilities, []string{fragment})
			continue
		}
		possibilities, err := getPossibleVariants(afterTrim, fragments)
		if err != nil {
			continue
		}
		for _, possibility := range possibilities {
			returnPossibilities = append(returnPossibilities, append(possibility, fragment))
		}
	}

	if nil == returnPossibilities {
		return nil, errors.New("not possible variant")
	}

	return returnPossibilities, nil
}
