package anagram

import (
	"sort"
	"strings"
)

// convert lowercase words to sorted slices of byte values and compare
func FindAnagrams(dictionary []string, word string) []string {
	// lowercase word and convert alphabetic characters to slice of ints
	cleanWordInts := getCleanIntSlice(word)
	// sort in preparation for comparison
	sort.Ints(cleanWordInts)

	// initialise result slice
	result := []string{}

	for index, dicWord := range dictionary {
		// lowercase dicWord and convert alphabetic characters to slice of ints
		cleanDicWordInts := getCleanIntSlice(dicWord)

		// eliminate if charcount different from word under exam
		if len(cleanWordInts) != len(cleanDicWordInts) {
			continue
		}

		// eliminate if dicWord is the same as word under exam
		if strings.EqualFold(strings.TrimSpace(word), strings.TrimSpace(dicWord)) {
			continue
		}

		// sort dictionary word for comparison
		sort.Ints(cleanDicWordInts)

		// compare and save if matched
		if areEqualIntSlices(cleanWordInts, cleanDicWordInts) {
			result = append(result, dictionary[index])
		}
	}

	return result
}

// Convert string to lowercase and make a slice of the integer values
// of the bytes that are in the range 'a' to 'z'
func getCleanIntSlice(s string) []int {
	// lowercase word
	s = strings.ToLower(s)

	// strip non alphanumeric and make slice of int values
	sBytes := []byte(s)
	cleanedInts := make([]int, 0, len(s))
	for _, b := range sBytes {
		if b > byte(96) && b < byte(123) {
			cleanedInts = append(cleanedInts, int(b))
		}
	}

	return cleanedInts
}

// check if 2 slices with the same length are equal
func areEqualIntSlices(word, dic []int) bool {
	for i := range word {
		if word[i] != dic[i] {
			return false
		}
	}
	return true
}
