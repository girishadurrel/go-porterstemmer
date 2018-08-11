package porterstemmer

func preprocess(s []rune) {
	normalizeApostrophes(s)
	s = trimLeftApostrophes(s)
}

func postprocess() {

}

//remove all the apostrophes 's or ' at the end
func step0(s []rune) []rune {
	suffixWithApos := []string{"'s", "'s'", "'"}

	for _, suffix := range suffixWithApos {
		if isSuffix, index := isSuffix(s, suffix); isSuffix {
			return s[:index]
		}
	}

	return s
}

func step1a(s []rune) []rune {

	// Initialize.
	var result []rune = s

	lenS := len(s)

	// Do it!
	if suffix := []rune("sses"); hasSuffix(s, suffix) {

		lenTrim := 2

		subSlice := s[:lenS-lenTrim]

		result = subSlice
	} else if contains, _ := hasSuffixes(s, [][]rune{[]rune("ies"), []rune("ied")}); contains {
		lenTrim := 1
		if lenS > 4 {
			lenTrim = 2
		}

		subSlice := s[:lenS-lenTrim]

		result = subSlice
	} else if conatins, _ := hasSuffixes(s, [][]rune{[]rune("ss"), []rune("us")}); conatins {

		result = s
	} else if suffix := []rune("s"); hasSuffix(s, suffix) {
		subSlice := s

		if containsVowel(s[:len(s)-2]) {
			lenSuffix := 1
			subSlice = s[:lenS-lenSuffix]
		}

		result = subSlice
	}

	// Return.
	return result
}

func step1b(s []rune) []rune {

	// Initialize.
	var result []rune = s

	lenS := len(s)
	r1Start, _ := getR1andR2Start(s)

	if contains, suffix := hasSuffixes(s, [][]rune{[]rune("eed"), []rune("eedly")}); contains {
		suffixLen := len(suffix)
		subSlice := s

		//suffix is contained within the r1Start
		if suffixLen <= lenS-r1Start {
			lenSuffix := 1
			if suffixLen == 5 { //"eedly"
				lenSuffix = 3
			}

			subSlice = s[:lenS-lenSuffix]
		}

		result = subSlice

	} else if contains, suffix := hasSuffixes(s, [][]rune{[]rune("ed"), []rune("edly"), []rune("ing"), []rune("ingly")}); contains {
		suffixLen := len(suffix)
		subSlice := s

		if containsVowel(s[:lenS-suffixLen]) {
			//delete the suffix

			s = s[:lenS-suffixLen]

			if contains, _ := hasSuffixes(s, [][]rune{[]rune("at"), []rune("bl"), []rune("iz")}); contains {

			}
		}

		result = subSlice
	}

	// Return.
	return result
}

func step1c(s []rune) []rune {

	// Initialize.
	lenS := len(s)

	result := s

	// Do it!
	if 2 > lenS {
		/////////// RETURN
		return result
	}

	if 'y' == s[lenS-1] && containsVowel(s[:lenS-1]) {

		result[lenS-1] = 'i'

	} else if 'Y' == s[lenS-1] && containsVowel(s[:lenS-1]) {

		result[lenS-1] = 'I'

	}

	// Return.
	return result
}

func step2(s []rune) []rune {

	// Initialize.
	lenS := len(s)

	result := s

	// Do it!
	if suffix := []rune("ational"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-5] = 'e'
			result = result[:lenS-4]
		}
	} else if suffix := []rune("tional"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = result[:lenS-2]
		}
	} else if suffix := []rune("enci"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-1] = 'e'
		}
	} else if suffix := []rune("anci"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-1] = 'e'
		}
	} else if suffix := []rune("izer"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-1]
		}
	} else if suffix := []rune("bli"); hasSuffix(s, suffix) { // --DEPARTURE--
		//		} else if suffix := []rune("abli") ; hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-1] = 'e'
		}
	} else if suffix := []rune("alli"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-2]
		}
	} else if suffix := []rune("entli"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-2]
		}
	} else if suffix := []rune("eli"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-2]
		}
	} else if suffix := []rune("ousli"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-2]
		}
	} else if suffix := []rune("ization"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-5] = 'e'

			result = s[:lenS-4]
		}
	} else if suffix := []rune("ation"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-3] = 'e'

			result = s[:lenS-2]
		}
	} else if suffix := []rune("ator"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-2] = 'e'

			result = s[:lenS-1]
		}
	} else if suffix := []rune("alism"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-3]
		}
	} else if suffix := []rune("iveness"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-4]
		}
	} else if suffix := []rune("fulness"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-4]
		}
	} else if suffix := []rune("ousness"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-4]
		}
	} else if suffix := []rune("aliti"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result = s[:lenS-3]
		}
	} else if suffix := []rune("iviti"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-3] = 'e'

			result = result[:lenS-2]
		}
	} else if suffix := []rune("biliti"); hasSuffix(s, suffix) {
		if 0 < measure(s[:lenS-len(suffix)]) {
			result[lenS-5] = 'l'
			result[lenS-4] = 'e'

			result = result[:lenS-3]
		}
	} else if suffix := []rune("logi"); hasSuffix(s, suffix) { // --DEPARTURE--
		if 0 < measure(s[:lenS-len(suffix)]) {
			lenTrim := 1

			result = s[:lenS-lenTrim]
		}
	}

	// Return.
	return result
}

func step3(s []rune) []rune {

	// Initialize.
	lenS := len(s)
	result := s

	// Do it!
	if suffix := []rune("icate"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		if 0 < measure(s[:lenS-lenSuffix]) {
			result = result[:lenS-3]
		}
	} else if suffix := []rune("ative"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 0 < m {
			result = subSlice
		}
	} else if suffix := []rune("alize"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		if 0 < measure(s[:lenS-lenSuffix]) {
			result = result[:lenS-3]
		}
	} else if suffix := []rune("iciti"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		if 0 < measure(s[:lenS-lenSuffix]) {
			result = result[:lenS-3]
		}
	} else if suffix := []rune("ical"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		if 0 < measure(s[:lenS-lenSuffix]) {
			result = result[:lenS-2]
		}
	} else if suffix := []rune("ful"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 0 < m {
			result = subSlice
		}
	} else if suffix := []rune("ness"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 0 < m {
			result = subSlice
		}
	}

	// Return.
	return result
}

func step4(s []rune) []rune {

	// Initialize.
	lenS := len(s)
	result := s

	// Do it!
	if suffix := []rune("al"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = result[:lenS-lenSuffix]
		}
	} else if suffix := []rune("ance"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = result[:lenS-lenSuffix]
		}
	} else if suffix := []rune("ence"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = result[:lenS-lenSuffix]
		}
	} else if suffix := []rune("er"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ic"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("able"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ible"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ant"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ement"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ment"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ent"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ion"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		c := subSlice[len(subSlice)-1]

		if 1 < m && ('s' == c || 't' == c) {
			result = subSlice
		}
	} else if suffix := []rune("ou"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ism"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ate"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("iti"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ous"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ive"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	} else if suffix := []rune("ize"); hasSuffix(s, suffix) {
		lenSuffix := len(suffix)

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	}

	// Return.
	return result
}

func step5a(s []rune) []rune {

	// Initialize.
	lenS := len(s)
	result := s

	// Do it!
	if 'e' == s[lenS-1] {
		lenSuffix := 1

		subSlice := s[:lenS-lenSuffix]
		if len(subSlice) == 0 {
			return result
		}
		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		} else if c := subSlice[len(subSlice)-1]; 1 == m && !(hasConsonantVowelConsonantSuffix(subSlice) && 'w' != c && 'x' != c && 'y' != c) {
			result = subSlice
		}
	}

	// Return.
	return result
}

func step5b(s []rune) []rune {

	// Initialize.
	lenS := len(s)
	result := s

	// Do it!
	if 2 < lenS && 'l' == s[lenS-2] && 'l' == s[lenS-1] {

		lenSuffix := 1

		subSlice := s[:lenS-lenSuffix]

		m := measure(subSlice)

		if 1 < m {
			result = subSlice
		}
	}

	// Return.
	return result
}
