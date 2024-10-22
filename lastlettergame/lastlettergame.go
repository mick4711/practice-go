package lastlettergame

import "fmt"

// make a struct with word, startCount, endCount
type wordParam struct {
	startRune  rune
	endRune    rune
	startCount int
	endCount   int
}

func Sequence(words []string) []string {
	// make a map of word params with start and end letters and joining word counts
	wordParams := getWordParams(words)
	for word, wordParam := range wordParams {
		fmt.Println(word, wordParam.startCount, wordParam.endCount)
	}

	// find middle with positive start and end counts
	for word, wordParam := range wordParams {
		fmt.Println(word, wordParam.startCount, wordParam.endCount)
	}
	// append valid start and end values
	return []string{"ab", "bc", "cd"}
}

// make a map of word params with start and end letters and joining word counts
func getWordParams(words []string) map[string]*wordParam {
	// initialise maps
	wordParams := make(map[string]*wordParam, len(words))
	startRunes := make(map[rune]int)
	endRunes := make(map[rune]int)

	// fill map with start and end letters keeping counts
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
	fmt.Println(startRunes, endRunes)

	// update map with start and end counts
	for _, wordParam := range wordParams {
		if count, ok := endRunes[wordParam.startRune]; ok {
			wordParam.endCount = count
		}
		if count, ok := startRunes[wordParam.endRune]; ok {
			wordParam.startCount = count
		}
	}

	return wordParams
}
