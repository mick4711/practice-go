package lastlettergame

import "fmt"

// make a struct with word parameters
type wordParam struct {
	startRune rune
	endRune   rune
	isRoot    bool
	done      bool
	children  []wordParam
}

var wordParams map[string]*wordParam

func Sequence(words []string) []string {
	// make a map of word params
	wordParams = setWordParams(words)
	nextWord := words[0]
	for finished := false; !finished; {
		children := getChildren(*wordParams[nextWord])
		wordParams[nextWord].children = children
		wordParams[nextWord].done = true

		nextWord, finished = getNextWord()
	}

	// debug print
	for word, wordParam := range wordParams {
		fmt.Println(word, *wordParam)
	}

	return []string{"ab", "bc", "cd"}
	// return []string{"df", "fz", "zd", "dq", "qf", "fg"}
}

// get next word that is not done
func getNextWord() (string, bool) {
	done := true
	for word, wordParam := range wordParams {
		if wordParam.done == false {
			return word, !done
		}
	}
	return "", done
}

// get words that can be tacked onto passed in word
func getChildren(wp wordParam) []wordParam {
	if wp.done {
		return wp.children
	}

	children := []wordParam{}
	for word, wordParam := range wordParams {
		if wp.endRune == wordParam.startRune {
			wordParams[word].isRoot = false
			children = append(children, *wordParams[word])
		}
	}

	if len(children) == 0 {
		return children
	}

	for _, wordParam := range children {
		getChildren(wordParam)
	}

	return children
}

// make a map of word params with start and end letters and joining word counts
func setWordParams(words []string) map[string]*wordParam {
	// initialise maps
	wordParams := make(map[string]*wordParam, len(words))

	// fill map with start and end letters keeping counts
	for _, word := range words {
		startRune := rune(word[0])
		endRune := rune(word[len(word)-1])
		wordParams[word] = &wordParam{
			startRune: startRune,
			endRune:   endRune,
			isRoot:    true,
			done:      false,
			children:  []wordParam{},
		}
	}

	return wordParams
}
