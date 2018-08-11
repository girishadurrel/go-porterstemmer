package porterstemmer

//replace all different types of apostrophes, with the code point 39 (')
func normalizeApostrophes(s []rune) {
	for index, r := range s {
		if r == 8217 || r == 8216 || r == 8219 {
			s[index] = 39
		}
	}
}

//remove until the first apostrophe is met
func trimLeftApostrophes(s []rune) []rune {
	index := 0
	r := rune(0)
	for index, r = range s {
		if r != 39 {
			break
		}
	}

	return s[index:]
}

func isSuffix(s []rune, suffix string) (bool, int) {
	suffixInRune := []rune(suffix)
	suffixIndex := -1
	totalLength := len(s)
	suffixLen := len(suffixInRune)

	if suffixLen > totalLength {
		return false, -1
	} else {
		suffixIndex = totalLength - suffixLen
		for index, item := range suffixInRune {
			if s[suffixIndex+index] != item {
				return false, -1
			}
		}
	}

	return true, suffixIndex
}

func isConsonant(s []rune, i int) bool {

	//DEBUG
	//log.Printf("isConsonant: [%+v]", string(s[i]))

	result := true

	switch s[i] {
	case 'a', 'e', 'i', 'o', 'u':
		result = false
	case 'y':
		if 0 == i {
			result = true
		} else {
			result = !isConsonant(s, i-1)
		}
	default:
		result = true
	}

	return result
}

func measure(s []rune) uint {

	// Initialize.
	lenS := len(s)
	result := uint(0)
	i := 0

	// Short Circuit.
	if 0 == lenS {
		/////////// RETURN
		return result
	}

	// Ignore (potential) consonant sequence at the beginning of word.
	for isConsonant(s, i) {

		//DEBUG
		//log.Printf("[measure([%s])] Eat Consonant [%d] -> [%s]", string(s), i, string(s[i]))

		i++
		if i >= lenS {
			/////////////// RETURN
			return result
		}
	}

	// For each pair of a vowel sequence followed by a consonant sequence, increment result.
Outer:
	for i < lenS {

		for !isConsonant(s, i) {

			//DEBUG
			//log.Printf("[measure([%s])] VOWEL [%d] -> [%s]", string(s), i, string(s[i]))

			i++
			if i >= lenS {
				/////////// BREAK
				break Outer
			}
		}
		for isConsonant(s, i) {

			//DEBUG
			//log.Printf("[measure([%s])] CONSONANT [%d] -> [%s]", string(s), i, string(s[i]))

			i++
			if i >= lenS {
				result++
				/////////// BREAK
				break Outer
			}
		}
		result++
	}

	// Return
	return result
}

func hasSuffixes(s []rune, suffixes [][]rune) (bool, []rune) {
	for _, suffix := range suffixes {
		if hasSuffix(s, suffix) {
			return true, suffix
		}
	}

	return false, nil
}

func hasSuffix(s, suffix []rune) bool {

	lenSMinusOne := len(s) - 1
	lenSuffixMinusOne := len(suffix) - 1

	if lenSMinusOne <= lenSuffixMinusOne {
		return false
	} else if s[lenSMinusOne] != suffix[lenSuffixMinusOne] { // I suspect checking this first should speed this function up in practice.
		/////// RETURN
		return false
	} else {

		for i := 0; i < lenSuffixMinusOne; i++ {

			if suffix[i] != s[lenSMinusOne-lenSuffixMinusOne+i] {
				/////////////// RETURN
				return false
			}

		}

	}

	return true
}

func containsVowel(s []rune) bool {

	lenS := len(s)

	for i := 0; i < lenS; i++ {

		if !isConsonant(s, i) {
			/////////// RETURN
			return true
		}

	}

	return false
}

func hasRepeatDoubleConsonantSuffix(s []rune) bool {

	// Initialize.
	lenS := len(s)

	result := false

	// Do it!
	if 2 > lenS {
		result = false
	} else if s[lenS-1] == s[lenS-2] && isConsonant(s, lenS-1) { // Will using isConsonant() cause a problem with "YY"?
		result = true
	} else {
		result = false
	}

	// Return,
	return result
}

func hasConsonantVowelConsonantSuffix(s []rune) bool {

	// Initialize.
	lenS := len(s)

	result := false

	// Do it!
	if 3 > lenS {
		result = false
	} else if isConsonant(s, lenS-3) && !isConsonant(s, lenS-2) && isConsonant(s, lenS-1) {
		result = true
	} else {
		result = false
	}

	// Return
	return result
}

func getConstIndexAfterVowel(s []rune) int {
	vowelIndex := -1
	startIndex := -1

	for index := range s {
		if !isConsonant(s, index) {
			vowelIndex = index
			break
		}
	}

	if vowelIndex != -1 {
		if len(s) > vowelIndex+1 {
			for index := range s[vowelIndex+1:] {
				if isConsonant(s, index+vowelIndex+1) {
					startIndex = index + vowelIndex + 1
					break
				}
			}
		}
	}

	return (startIndex + 1) //plus one because region starts after :)
}

func getR1andR2Start(s []rune) (int, int) {
	//region after the first non-vowel after the first vowel
	r1Start := len(string(s))
	//region after the first non-vowel following the first vowel in r1
	r2Start := len(string(s))

	if returnVal := getConstIndexAfterVowel(s); returnVal > -1 {
		r1Start = returnVal
	}

	if r1Start+1 < len(s) {
		if returnVal := getConstIndexAfterVowel(s[r1Start:]); returnVal > -1 {
			r2Start = returnVal + r1Start
		}
	}

	return r1Start, r2Start
}
