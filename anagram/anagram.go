package anagram

import "strings"

func FindAnagrams(dictionary []string, word string) (result []string) {

	length := len(strings.ReplaceAll(word, " ", ""))
	result = []string{}
	word = strings.ToLower(word)
	for _, dictionaryWord := range dictionary {
		tmpDictionaryWord := strings.ToLower(dictionaryWord)
		if tmpDictionaryWord == word {
			continue
		}
		if length == len(strings.ReplaceAll(tmpDictionaryWord, " ", "")) && 0 < len(strings.ReplaceAll(tmpDictionaryWord, " ", "")) {
			if isSameLetters(word, tmpDictionaryWord) {
				result = append(result, dictionaryWord)
			}
		}
	}

	return result
}

func isSameLetters(word string, dictionaryWord string) bool {
	processedWord := strings.ReplaceAll(strings.ToLower(word), " ", "")
	processedDictionaryWord := strings.ReplaceAll(strings.ToLower(dictionaryWord), " ", "")
	for _, c := range word {
		processedWord = strings.ReplaceAll(processedWord, string(c), "")
		processedDictionaryWord = strings.ReplaceAll(processedDictionaryWord, string(c), "")
		if len(processedWord) != len(processedDictionaryWord) {
			return false
		}
	}
	return true
}
