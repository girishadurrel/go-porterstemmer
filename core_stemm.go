package porterstemmer

func preprocess(s []rune) []rune {
	s = normalizeApostrophes(s)
	s = trimLeftApostrophes(s)

	return s
}

//remove all the apostrophes 's or ' at the end
func step0(s []rune, r1Start, r2Start int) ([]rune, int, int) {
	lenS := len(s)
	result := s

	suffixes := [][]rune{[]rune("'s'"), []rune("'s"), []rune("'")}

	if contains, suffix := hasSuffixes(s, suffixes); contains {
		suffixLen := len(suffix)

		s = s[:lenS-suffixLen]

		r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
		result = s
	}

	return result, r1Start, r2Start
}

func step1a(s []rune, r1Start, r2Start int) ([]rune, int, int) {
	// Initialize.
	lenS := len(s)
	result := s

	suffixes := [][]rune{
		[]rune("sses"), []rune("ied"), []rune("ies"), []rune("us"), []rune("ss"),
		[]rune("s"),
	}

	if contains, suffix := hasSuffixes(s, suffixes); contains {
		suffixLen := len(suffix)
		suffixInString := string(suffix)

		if suffixInString == "sses" {
			lenWithoutSuffix := lenS - suffixLen
			s = append(s[:lenWithoutSuffix], []rune("ss")...)

			r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
			result = s
		} else if suffixInString == "ies" || suffixInString == "ied" {
			lenWithoutSuffix := lenS - suffixLen
			replaceRune := []rune("ie")

			if lenS > 4 {
				replaceRune = []rune("i")
			}

			s = append(s[:lenWithoutSuffix], replaceRune...)
			r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
			result = s

		} else if suffixInString == "ss" || suffixInString == "us" {
			result = s
		} else if suffixInString == "s" {
			for i := 0; i < lenS-2; i++ {
				if !isConsonant(s, i) {
					s = s[:lenS-suffixLen]

					r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
					result = s

				}
			}
		}
	}

	// Return.
	return result, r1Start, r2Start
}

func step1b(s []rune, r1Start, r2Start int) ([]rune, int, int) {
	// Initialize.
	lenS := len(s)
	result := s

	orgR1Start := r1Start
	orgR2Start := r2Start

	suffixes := [][]rune{
		[]rune("eed"), []rune("eedly"), []rune("ed"), []rune("edly"), []rune("ing"), []rune("ingly"),
	}

	if contains, suffix := hasSuffixes(s, suffixes); contains {
		suffixLen := len(suffix)
		suffixInString := string(suffix)

		if suffixInString == "eed" || suffixInString == "eedly" {
			if suffixLen <= lenS-r1Start {
				lenWithoutSuffix := lenS - suffixLen

				s = append(s[:lenWithoutSuffix], []rune("ee")...)

				r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
				result = s
			}
		} else if suffixInString == "ed" || suffixInString == "edly" || suffixInString == "ing" || suffixInString == "ingly" {
			lenWithoutSuffix := lenS - suffixLen

			if containsVowel(s[:lenWithoutSuffix]) {
				s = s[:lenWithoutSuffix]

				suffixes := [][]rune{
					[]rune("at"), []rune("bl"), []rune("iz"), []rune("bb"), []rune("dd"), []rune("ff"), []rune("gg"),
					[]rune("mm"), []rune("nn"), []rune("pp"), []rune("rr"), []rune("tt"),
				}

				if contains, suffix := hasSuffixes(s, suffixes); contains {
					suffixLen := len(suffix)
					lenWithoutSuffix := len(s) - suffixLen
					suffixInString := string(suffix)

					if suffixInString == "at" || suffixInString == "bl" || suffixInString == "iz" {
						s = append(s[:lenWithoutSuffix], []rune(suffixInString+"e")...)

						r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
						result = s

					} else if suffixInString == "bb" || suffixInString == "dd" || suffixInString == "ff" || suffixInString == "gg" || suffixInString == "mm" || suffixInString == "nn" || suffixInString == "pp" || suffixInString == "rr" || suffixInString == "tt" {
						s = s[:len(s)-1]

						r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
						result = s
					}
				} else {
					if isShortWord(s, r1Start) {
						s = append(s, []rune("e")...)
						result = s

						r1Start = len(s)
						r2Start = len(s)

						return result, r1Start, r2Start
					} else {
						result = s
					}
				}

				currLen := len(s)

				if orgR1Start < currLen {
					r1Start = orgR1Start
				} else {
					r1Start = currLen
				}

				if orgR2Start < currLen {
					r2Start = orgR2Start
				} else {
					r2Start = currLen
				}
			}
		}
	}

	// Return.
	return result, r1Start, r2Start
}

func step1c(s []rune, r1Start, r2Start int) ([]rune, int, int) {
	// Initialize.
	lenS := len(s)
	result := s

	//if the last letter is a y, then replace it with a 'i'
	if lenS > 2 && s[lenS-1] == 121 && isConsonant(s, lenS-2) {
		s[lenS-1] = 105
		result = s
	}

	// Return.
	return result, r1Start, r2Start
}

func step2(s []rune, r1Start, r2Start int) ([]rune, int, int) {
	// Initialize.
	lenS := len(s)
	result := s

	suffixes := [][]rune{
		[]rune("ational"), []rune("fulness"), []rune("iveness"), []rune("ization"), []rune("ousness"),
		[]rune("biliti"), []rune("lessli"), []rune("tional"), []rune("alism"), []rune("aliti"), []rune("ation"),
		[]rune("entli"), []rune("fulli"), []rune("iviti"), []rune("ousli"), []rune("anci"), []rune("abli"),
		[]rune("alli"), []rune("ator"), []rune("enci"), []rune("izer"), []rune("bli"), []rune("ogi"), []rune("li"),
	}

	replaceMap := map[string]string{
		"tional":  "tion",
		"enci":    "ence",
		"anci":    "ance",
		"abli":    "able",
		"entli":   "ent",
		"izer":    "ize",
		"ization": "ize",
		"ational": "ate",
		"ation":   "ate",
		"ator":    "ate",
		"alism":   "al",
		"aliti":   "al",
		"alli":    "al",
		"fulness": "ful",
		"ousli":   "ous",
		"ousness": "ous",
		"iveness": "ive",
		"iviti":   "ive",
		"biliti":  "ble",
		"bli":     "ble",
		"fulli":   "ful",
		"lessli":  "less",
	}

	liChs := map[rune]bool{
		99: true, 100: true, 101: true, 103: true, 104: true, 107: true, 109: true, 110: true, 114: true, 116: true,
	}

	if contains, suffix := hasSuffixes(s, suffixes); contains {
		suffixLen := len(suffix)
		suffixInString := string(suffix)

		if suffixLen <= lenS-r1Start {
			if suffixInString == "li" {
				if lenS >= 3 {
					if _, ok := liChs[s[lenS-3]]; ok {
						s = s[:lenS-suffixLen]

						r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
						result = s
					}
				}

			} else if suffixInString == "ogi" {
				if lenS >= 4 && s[lenS-4] == 108 {
					lenWithoutSuffix := lenS - suffixLen
					s = append(s[:lenWithoutSuffix], []rune("og")...)

					r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
					result = s
				}

			} else {
				if replacementSuffix, ok := replaceMap[suffixInString]; ok {
					replacementSuffixInRune := []rune(replacementSuffix)
					lenWithoutSuffix := lenS - suffixLen
					s = append(s[:lenWithoutSuffix], replacementSuffixInRune...)

					r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
					result = s
				}

			}
		}
	}

	// Return.
	return result, r1Start, r2Start
}

func step3(s []rune, r1Start, r2Start int) ([]rune, int, int) {
	// Initialize.
	lenS := len(s)
	result := s

	suffixes := [][]rune{
		[]rune("ational"), []rune("tional"), []rune("alize"), []rune("icate"),
		[]rune("ative"), []rune("iciti"), []rune("ical"), []rune("ful"),
		[]rune("ness"),
	}

	replaceMap := map[string]string{
		"ational": "ate",
		"tional":  "tion",
		"alize":   "al",
		"icate":   "ic",
		"iciti":   "ic",
		"ical":    "ic",
	}

	if contains, suffix := hasSuffixes(s, suffixes); contains {
		suffixLen := len(suffix)
		suffixInString := string(suffix)

		if suffixLen <= lenS-r1Start {
			if suffixInString == "ative" {
				if lenS-r2Start >= suffixLen {
					s = s[:lenS-suffixLen]

					r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
					result = s
				}
			} else if suffixInString == "ful" || suffixInString == "ness" {
				lenWithoutSuffix := lenS - suffixLen
				s = s[:lenWithoutSuffix]

				r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
				result = s
			} else {
				if replacementSuffix, ok := replaceMap[suffixInString]; ok {
					replacementSuffixInRune := []rune(replacementSuffix)
					lenWithoutSuffix := lenS - suffixLen

					s = append(s[:lenWithoutSuffix], replacementSuffixInRune...)
					r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
					result = s
				}
			}
		}
	}

	// Return.
	return result, r1Start, r2Start
}

func step4(s []rune, r1Start, r2Start int) ([]rune, int, int) {
	// Initialize.
	lenS := len(s)
	result := s

	suffixes := [][]rune{
		[]rune("ement"), []rune("ance"), []rune("ence"), []rune("able"), []rune("ible"), []rune("ment"),
		[]rune("ent"), []rune("ant"), []rune("ism"), []rune("ate"), []rune("iti"), []rune("ous"), []rune("ive"),
		[]rune("ize"), []rune("ion"), []rune("al"), []rune("er"), []rune("ic"),
	}

	if contains, suffix := hasSuffixes(s, suffixes); contains {
		suffixLen := len(suffix)
		suffixInString := string(suffix)

		if suffixLen <= lenS-r2Start {
			if suffixInString == "ion" {
				if lenS >= 4 {
					if s[lenS-4] == 115 || s[lenS-4] == 116 {
						s := s[:lenS-suffixLen]

						r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
						result = s
					}
				}
			} else {
				s := s[:lenS-suffixLen]

				r1Start, r2Start = updateR1R2(len(s), r1Start, r2Start)
				result = s
			}
		}
	}

	// Return.
	return result, r1Start, r2Start
}

func step5(s []rune, r1Start, r2Start int) ([]rune, int, int) {
	// Initialize.
	lenS := len(s)
	result := s

	if r1Start <= lenS-1 {
		if s[lenS-1] == 101 { //ends with an 'e'
			if r2Start <= lenS-1 || !endsShortSyllable(s, lenS-1) {
				s = s[:lenS-1]

				result = s
			} else if r2Start <= lenS-1 && s[lenS-1] == 108 && lenS-2 >= 0 && s[lenS-2] == 108 {
				s = s[:lenS-1]

				result = s
			}
		}
	}

	// Return.
	return result, r1Start, r2Start
}
