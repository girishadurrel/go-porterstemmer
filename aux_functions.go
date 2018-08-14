package porterstemmer

//replace all different types of apostrophes, with the code point 39 (')
func normalizeApostrophes(s []rune) []rune {
	for index, r := range s {
		if r == 8217 || r == 8216 || r == 8219 {
			s[index] = 39
		}
	}

	return s
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

func isShortWord(s []rune, r1Start int) bool {
	if r1Start < len(s) {
		return false
	}

	return endsShortSyllable(s, len(s))
}

func endsShortSyllable(s []rune, index int) bool {
	if index == 2 {
		//starts with a vowel and ends with a consonent, then its true
		if !isConsonant(s, 0) && isConsonant(s, 1) {
			return true
		} else {
			return false
		}
	} else if index >= 3 {
		s1 := s[index-1]

		//check if s1 is precceded by a vowel is already checked
		if isConsonant(s, index-1) && s1 != 119 && s1 != 120 && !isConsonant(s, index-2) && isConsonant(s, index-3) {
			if s1 == 121 && !isConsonant(s, index-2) { //check if y is preceeded by a vowel
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	}

	return false
}

func updateR1R2(lenS, r1Start, r2Start int) (int, int) {
	if r1Start > lenS {
		r1Start = lenS
	}

	if r2Start > lenS {
		r2Start = lenS
	}

	return r1Start, r2Start
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

func hasSuffixes(s []rune, suffixes [][]rune) (bool, []rune) {
	for _, suffix := range suffixes {
		if hasSuffix(s, suffix) {
			return true, suffix
		}
	}

	return false, nil
}

func hasSuffix(s, suffix []rune) bool {
	lenS := len(s)
	suffixLen := len(suffix)

	if lenS < suffixLen {
		return false
	} else {
		for i := 0; i < suffixLen; i++ {
			if s[lenS-1-i] != suffix[suffixLen-1-i] {
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

//need stat index to check on cases of 'y'
//'y' before a consonent and 'y' before a vowel behaves differently
func getConstIndexAfterVowel(s []rune, startIndex int) int {
	vowelIndex := -1
	returnVal := -1

	for index := range s[startIndex:] {
		if !isConsonant(s, index+startIndex) {
			vowelIndex = index + startIndex
			break
		}
	}

	if vowelIndex != -1 {
		if len(s) > vowelIndex+1 {
			for index := range s[vowelIndex+1:] {
				if isConsonant(s, index+vowelIndex+1) {
					returnVal = index + vowelIndex + 1
					break
				}
			}
		}
	}

	if returnVal != -1 && vowelIndex != -1 { //index has been set!!!
		return (returnVal + 1) //plus one because region starts after :)
	} else {
		return -1
	}
}

func getR1andR2Start(s []rune) (int, int) {
	//region after the first non-vowel after the first vowel
	r1Start := len(string(s))
	//region after the first non-vowel following the first vowel in r1
	r2Start := len(string(s))

	prefixes := [][]rune{
		[]rune("gener"), []rune("commun"), []rune("arsen"),
	}

	if contains, prefix := hasPrefixes(s, prefixes); contains {
		r1Start = len(prefix)
	} else if returnVal := getConstIndexAfterVowel(s, 0); returnVal > -1 {
		r1Start = returnVal
	}

	if r1Start+1 < len(s) {
		if returnVal := getConstIndexAfterVowel(s, r1Start); returnVal > -1 {
			r2Start = returnVal
		}
	}

	return r1Start, r2Start
}

func hasPrefixes(s []rune, prefixes [][]rune) (bool, []rune) {
	for _, prefix := range prefixes {
		if hasPrefix(s, prefix) {
			return true, prefix
		}
	}

	return false, nil
}

func hasPrefix(s []rune, prefix []rune) bool {
	lenS := len(s)
	prefixLen := len(prefix)

	if prefixLen > lenS {
		return false
	} else {
		for i := 0; i < prefixLen; i++ {
			if s[i] != prefix[i] {
				return false
			}
		}
	}

	return true
}
