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

func isShortWord(s []rune, r1Start int) bool {
	if r1Start < len(s) || endsShortSyllable(s, len(s)) {
		return true
	}

	return false
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
			return true
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
