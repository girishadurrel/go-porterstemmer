package porterstemmer

import (
	"fmt"
	"unicode"
)

type PorterStemmer struct {
	debug     bool
	aggrStem  bool
	ignoreMap map[string]bool
}

func (e *PorterStemmer) Init(debug, aggrStem bool, ignoreMap map[string]bool) {
	e.debug = debug
	e.aggrStem = aggrStem
	e.ignoreMap = ignoreMap
}

func (e *PorterStemmer) StemString(s string) (string, []string) {

	// Convert string to []rune
	runeArr := []rune(s)

	// Stem.
	runeArr, debugDetails := e.Stem(runeArr)

	// Convert []rune to string
	str := string(runeArr)

	// Return.
	return str, debugDetails
}

func (e *PorterStemmer) Stem(s []rune) ([]rune, []string) {

	// Initialize.
	lenS := len(s)
	debugDetails := make([]string, 0)

	_, doNotStemm := e.ignoreMap[string(s)]

	// Short circuit.
	if 0 == lenS || doNotStemm {
		/////////// RETURN
		debugDetails = append(debugDetails, fmt.Sprintf("input: %s is in the ignore dict", string(s)))
		return s, debugDetails
	}

	// Make all runes lowercase.
	for i := 0; i < lenS; i++ {
		s[i] = unicode.ToLower(s[i])
	}

	// Stem
	result, debugDetails := e.StemWithoutLowerCasing(s)

	// Return.
	return result, debugDetails
}

func (e *PorterStemmer) StemWithoutLowerCasing(s []rune) ([]rune, []string) {

	// Initialize.
	lenS := len(s)
	debugLines := make([]string, 0)

	// Words that are of length 2 or less is already stemmed.
	// Don't do anything.
	if 2 >= lenS {
		/////////// RETURN
		debugLines = append(debugLines, "input less than 2 runs, return")
		return s, debugLines
	}

	// Stem
	if e.debug {
		debugLines = append(debugLines, fmt.Sprintf("input: %s", string(s)))
	}
	s = step0(s)  // remove all apostrophes
	s = step1a(s) //endings of nouns (step5a,b is also important for nouns)
	if e.debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (1a): %s", string(s)))
	}

	if !e.aggrStem {
		s = step1b(s) //verb endings that usually end with (ed, ing)
		s = step1c(s) //anything that ends in 'y', replace with 'i' verbs (adverb)
		//example play -> plai, lay -> lai etc etc
		if e.debug {
			debugLines = append(debugLines, fmt.Sprintf("after step (1b&c): %s", string(s)))
		}

		s = step2(s)
		if e.debug {
			debugLines = append(debugLines, fmt.Sprintf("after step (2): %s", string(s)))
		}

		s = step3(s)
		if e.debug {
			debugLines = append(debugLines, fmt.Sprintf("after step (3): %s", string(s)))
		}

		s = step4(s)
		if e.debug {
			debugLines = append(debugLines, fmt.Sprintf("after step (4): %s", string(s)))
		}
	}

	s = step5a(s)
	s = step5b(s)
	if e.debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (5): %s", string(s)))
	}

	// Return.
	return s, debugLines
}
