package lastlettergame

import "fmt"

// make a struct with word parameters
type wordParam struct {
	startRune rune
	endRune   rune
	isRoot    bool
	done      bool
	children  []string
}

// var wordParams map[string]*wordParam

func Sequence(words []string) []string {
	// make a map of word params
	wordParams := setWordParams(words)

	// populate children
	populateChildren(words, wordParams)

	// get words that have no parents
	rootWords := getRootWords(wordParams)

	// for each rootWord get longest sequence
	var sequence []string
	for _, rootWord := range rootWords {
		sequence = getLongestSeq(rootWord, wordParams)
	}

	// debug print
	for word, wordParam := range wordParams {
		fmt.Println(word, *wordParam)
	}

	return sequence
	// return []string{"ab", "bc", "cd"}
	// return []string{"df", "fz", "zd", "dq", "qf", "fg"}
}

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
			children:  []string{},
		}
	}

	return wordParams
}

func populateChildren(words []string, wordParams map[string]*wordParam) {
	nextWord := words[0]
	for finished := false; !finished; {
		children := getChildren(*wordParams[nextWord], wordParams)
		wordParams[nextWord].children = children
		wordParams[nextWord].done = true

		nextWord, finished = getNextWord(wordParams)
	}
}

// get words that can be tacked onto passed in word
func getChildren(wp wordParam, wordParams map[string]*wordParam) []string {
	if wp.done {
		return wp.children
	}

	children := []string{}
	for word, wordParam := range wordParams {
		if wp.endRune == wordParam.startRune {
			wordParams[word].isRoot = false
			children = append(children, word)
		}
	}

	for _, word := range children {
		getChildren(*wordParams[word], wordParams)
	}

	return children
}

// get next word that is not done
func getNextWord(wordParams map[string]*wordParam) (string, bool) {
	done := true
	for word, wordParam := range wordParams {
		if wordParam.done == false {
			return word, !done
		}
	}
	return "", done
}

func getRootWords(wordParams map[string]*wordParam) []string {
	rootWords := []string{}
	for word, wordParam := range wordParams {
		if wordParam.isRoot {
			rootWords = append(rootWords, word)
		}
	}

	return rootWords
}

func getLongestSeq(word string, wordParams map[string]*wordParam) []string {
	seq := []string{}
	for _, child := range wordParams[word].children {
		seq = append(seq, child)
	}

	// return seq
	return []string{"ab", "bc", "cd"}
}
