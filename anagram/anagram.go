package anagram

import (
	// "fmt"
	"sort"
	"strings"
)

func FindAnagrams(dictionary []string, word string) []string {
	// clean word
	cleanWordInts := cleanInts(word)
	sort.Ints(cleanWordInts)
	// fmt.Printf("clean word string %v\n", cleanWordInts)

	res := []string{}

	for index, dicWord := range dictionary {
		// eliminate if lettercount different
		cleanDicInts := cleanInts(dicWord)
		if len(cleanWordInts) != len(cleanDicInts) {
			continue
		}

		//TODO cache cleanDicInts keyed by index
		
		// eliminate if same word
		if strings.ToLower(strings.TrimSpace(word)) == strings.ToLower(strings.TrimSpace(dicWord)) {
			continue
		}

		sort.Ints(cleanDicInts)
		// fmt.Printf("possible dictionary word(%v) %v\n", index, cleanDicInts)
		if areEqual(cleanWordInts, cleanDicInts) {
			res = append(res, dictionary[index])
		}
	}

	return res
}

func cleanInts(s string) []int {
	// lowercase word
	s = strings.ToLower(s)

	// strip non alphanumeric
	sBytes := []byte(s)
	cleanedInts := make([]int, 0, len(s))
	for _, b := range sBytes {
		if b > byte(96) && b < byte(123) {
			cleanedInts = append(cleanedInts, int(b))
		}
	}

	return cleanedInts
}

func areEqual(word, dic []int) bool {
	for i := range word {
		if word[i] != dic[i] {
			return false
		}
	}
	return true
}
