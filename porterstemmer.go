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

	if lenS == 0 || doNotStemm {
		debugDetails = append(debugDetails, fmt.Sprintf("input: %s is in the ignore dict", string(s)))
		return s, debugDetails
	}

	for i := 0; i < lenS; i++ {
		s[i] = unicode.ToLower(s[i])
	}

	result, debugDetails := stemEngine(s, e.debug)

	return result, debugDetails
}

func stemEngine(s []rune, debug bool) ([]rune, []string) {
	// Initialize.
	lenS := len(s)
	debugLines := make([]string, 0)

	// Words that are of length 2 or less is already stemmed.
	// Don't do anything.
	if lenS <= 2 {
		/////////// RETURN
		debugLines = append(debugLines, "input less than 2 runs, return")
		return s, debugLines
	}

	// Stem
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("input: %s", string(s)))
	}

	r1Start, r2Start := getR1andR2Start(s)
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("r1Start: %d r2Start: %d", r1Start, r2Start))
	}

	s, r1Start, r2Start = step0(s, r1Start, r2Start) // remove all apostrophes
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (0): %s, r1Start: %d r2Start: %d", string(s), r1Start, r2Start))
	}

	s, r1Start, r2Start = step1a(s, r1Start, r2Start)
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (1a): %s r1Start: %d r2Start: %d", string(s), r1Start, r2Start))
	}

	s, r1Start, r2Start = step1b(s, r1Start, r2Start)
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (1b): %s r1Start: %d r2Start: %d", string(s), r1Start, r2Start))
	}

	s, r1Start, r2Start = step1c(s, r1Start, r2Start)
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (1c): %s r1Start: %d r2Start: %d", string(s), r1Start, r2Start))
	}

	s, r1Start, r2Start = step2(s, r1Start, r2Start)
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (2): %s r1Start: %d r2Start: %d", string(s), r1Start, r2Start))
	}

	s, r1Start, r2Start = step3(s, r1Start, r2Start)
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (3): %s r1Start: %d r2Start: %d", string(s), r1Start, r2Start))
	}

	s, r1Start, r2Start = step4(s, r1Start, r2Start)
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (4): %s r1Start: %d r2Start: %d", string(s), r1Start, r2Start))
	}

	s, r1Start, r2Start = step5(s, r1Start, r2Start)
	if debug {
		debugLines = append(debugLines, fmt.Sprintf("after step (5): %s r1Start: %d r2Start: %d", string(s), r1Start, r2Start))
	}

	// Return.
	return s, debugLines
}
