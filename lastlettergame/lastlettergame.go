package lastlettergame

import "fmt"

// make a struct with word, startCount, endCount
type wordParam struct {
	startRune  rune
	endRune    rune
	firstCount int
	lastCount  int
}

func Sequence(words []string) []string {
	wordParams := make(map[string]*wordParam, len(words))
	// iter words check end letter count at start of words
	startRunes := make(map[rune]int)
	endRunes := make(map[rune]int)
	for _, word := range words {
		startRune := rune(word[0])
		startRunes[startRune]++
		endRune := rune(word[len(word)-1])
		endRunes[endRune]++
		wordParams[word] = &wordParam{
			startRune,
			endRune,
			0,
			0,
		}
	}
	for word, wordParam := range wordParams {
		if count, ok := endRunes[wordParam.startRune]; ok {
			wordParams[word].lastCount = count
		}
		if count, ok := startRunes[wordParam.endRune]; ok {
			wordParam.firstCount = count
		}
	}
	for word, wordParam := range wordParams {
		fmt.Println(word, *&wordParam.firstCount, *&wordParam.lastCount)
	}

	fmt.Println(startRunes)
	fmt.Println(endRunes)

	// and start letter as count of word endings
	// find middle with positive start and end counts
	// append valid start and end values
	return []string{"ab", "bc", "cd"}
}
