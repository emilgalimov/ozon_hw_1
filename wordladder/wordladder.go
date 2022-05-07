package wordladder

import "errors"

func WordLadder(from string, to string, dic []string) int {

	if 0 == len(dic) {
		return 0
	}

	routes, err := possibleRoutes(from, to, dic)

	if nil != err {
		return 0
	}

	var minRoute int

	for _, route := range routes {
		length := len(route) + 1
		if 0 == minRoute || minRoute > length {
			minRoute = length
		}
	}

	return minRoute
}

func possibleRoutes(current string, needed string, dic []string) ([][]string, error) {

	if current == needed {
		return nil, nil
	}

	if 0 == len(dic) {
		return nil, errors.New("not found")
	}

	var result [][]string

	for i, word := range dic {
		if !isOneDifference(word, current) {
			continue
		}
		newDict := removeIndex(dic, i)
		newResultItems, err := possibleRoutes(word, needed, newDict)

		if nil != err {
			continue
		}

		var resultItem []string
		resultItem = append(resultItem, word)

		if len(newResultItems) == 0 {
			result = append(result, resultItem)
		}

		for _, newResultItem := range newResultItems {
			result = append(result, append(resultItem, newResultItem...))
		}
	}

	if len(result) == 0 {
		return nil, errors.New("not found")
	}

	return result, nil
}

func isOneDifference(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var count int

	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}

func removeIndex(s []string, index int) []string {
	newS := make([]string, len(s))
	copy(newS, s)
	return append(newS[:index], newS[index+1:]...)
}
