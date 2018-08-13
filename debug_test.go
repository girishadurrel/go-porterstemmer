package porterstemmer

import (
	"testing"
)

func Test_normalizeApostrophes(t *testing.T) {
	variants := [...]string{
		"\u2019xxx\u2019",
		"\u2018xxx\u2018",
		"\u201Bxxx\u201B",
		"’xxx’",
		"‘xxx‘",
		"‛xxx‛",
	}
	for _, v := range variants {
		output := normalizeApostrophes([]rune(v))
		if string(output) != "'xxx'" {
			t.Errorf("Expected \"'xxx'\", not \"%s\"", string(output))
		}
	}
}

func Test_preprocess(t *testing.T) {
	var wordTests = []struct {
		in  string
		out string
	}{
		{"arguing", "arguing"},
		{"'catty", "catty"},
		{"kyle’s", "kyle's"},
		//{"toy", "toY"},
	}
	for _, wt := range wordTests {
		output := preprocess([]rune(wt.in))
		if string(output) != wt.out {
			t.Errorf("Expected \"%s\", not \"%s\"", wt.out, string(output))
		}
	}
}

func Test_getR1andR2Start(t *testing.T) {
	var wordTests = []struct {
		word string
		r1   string
		r2   string
	}{
		{"crepuscular", "uscular", "cular"},
		{"beautiful", "iful", "ul"},
		{"beauty", "y", ""},
		{"eucharist", "harist", "ist"},
		{"animadversion", "imadversion", "adversion"},
		{"mistresses", "tresses", "ses"},
		{"sprinkled", "kled", ""},
		// Special cases below
		{"communism", "ism", "m"},
		{"arsenal", "al", ""},
		{"generalities", "alities", "ities"},
		{"embed", "bed", ""},
	}

	for _, testCase := range wordTests {
		r1Start, r2Start := getR1andR2Start([]rune(testCase.word))

		r1Str := testCase.word[r1Start:]
		r2Str := testCase.word[r2Start:]

		if r1Str != testCase.r1 || r2Str != testCase.r2 {
			t.Errorf("Expected \"{%s, %s}\", but got \"{%s, %s}\"", testCase.r1, testCase.r2, r1Str, r2Str)
		}
	}

}

func Test_isShortWord(t *testing.T) {
	var testCases = []struct {
		word    string
		isShort bool
	}{
		{"bed", true},
		{"shed", true},
		{"shred", true},
		{"bead", false},
		{"embed", false},
		{"beds", false},
	}
	for _, testCase := range testCases {
		s := []rune(testCase.word)
		r1Start, _ := getR1andR2Start(s)

		isShort := isShortWord(s, r1Start)
		if isShort != testCase.isShort {
			t.Errorf("Expected %t, but got %t for \"{%s, %s}\"", testCase.isShort, isShort, testCase.word, string(s[r1Start:]))
		}
	}
}

func Test_endsShortSyllable(t *testing.T) {
	var testCases = []struct {
		word   string
		pos    int
		result bool
	}{
		{"absolute", 7, true},
		{"ape", 2, true},
		{"rap", 3, true},
		{"trap", 4, true},
		{"entrap", 6, true},
		{"uproot", 6, false},
		{"bestow", 6, false},
		{"disturb", 7, false},
	}
	for _, testCase := range testCases {
		s := []rune(testCase.word)
		result := endsShortSyllable(s, testCase.pos)
		if result != testCase.result {
			t.Errorf("Expected endsShortSyllable(%s, %d) to return %t, not %t", testCase.word, testCase.pos, testCase.result, result)
		}
	}
}

type stepFunc func([]rune, int, int) ([]rune, int, int)

type testCase struct {
	inputStr string
	r1Start  int
	r2Start  int
	outStr   string
	outR1Str string
	outR2Str string
}

func runStepTest(t *testing.T, testCases []testCase, f stepFunc) {
	for _, testCase := range testCases {
		output, outR1, outR2 := f([]rune(testCase.inputStr), testCase.r1Start, testCase.r2Start)

		if string(output) != testCase.outStr || testCase.outR1Str != string(output[outR1:]) || testCase.outR2Str != string(output[outR2:]) {
			t.Errorf("Expected \"{%s, %s, %s}\", but got \"{%s, %s, %s}\"", testCase.outStr, testCase.outR1Str, testCase.outR2Str, string(output), string(output[outR1:]), string(output[outR2:]))
		}
	}
}

func Test_step0(t *testing.T) {
	var testCases = []testCase{
		{"general's", 5, 9, "general", "al", ""},
		{"general's'", 5, 10, "general", "al", ""},
		{"spices'", 4, 7, "spices", "es", ""},
	}
	runStepTest(t, testCases, step0)
}

func Test_step1a(t *testing.T) {
	var testCases = []testCase{
		{"ties", 0, 0, "tie", "tie", "tie"},
		{"cries", 0, 0, "cri", "cri", "cri"},
		{"mistresses", 3, 7, "mistress", "tress", "s"},
		{"ied", 3, 3, "ie", "", ""},
	}
	runStepTest(t, testCases, step1a)
}

func Test_step1b(t *testing.T) {
	var testCases = []testCase{
		{"exxeedly", 1, 8, "exxee", "xxee", ""},
		{"exxeed", 1, 7, "exxee", "xxee", ""},
		{"luxuriated", 3, 5, "luxuriate", "uriate", "iate"},
		{"luxuribled", 3, 5, "luxurible", "urible", "ible"},
		{"luxuriized", 3, 5, "luxuriize", "uriize", "iize"},
		{"luxuriedly", 3, 5, "luxuri", "uri", "i"},
		{"vetted", 3, 6, "vet", "", ""},
		{"hopping", 3, 7, "hop", "", ""},
		{"breed", 5, 5, "breed", "", ""},
		{"skating", 4, 6, "skate", "e", ""},
	}
	runStepTest(t, testCases, step1b)
}

func Test_step1c(t *testing.T) {
	var testCases = []testCase{
		{"cry", 3, 3, "cri", "", ""},
		{"say", 3, 3, "say", "", ""},
		{"by", 2, 2, "by", "", ""},
		{"xexby", 2, 5, "xexbi", "xbi", ""},
	}
	runStepTest(t, testCases, step1c)
}

func Test_step2(t *testing.T) {
	// Here I've faked R1 & R2 for simplicity
	var testCases = []testCase{
		{"fluentli", 5, 8, "fluentli", "tli", ""},
		// Test "tional"
		{"xxxtional", 3, 5, "xxxtion", "tion", "on"},
		// Test when "tional" doesn't fit in R1
		{"xxxtional", 4, 5, "xxxtional", "ional", "onal"},
		// Test "li"
		{"xxxcli", 3, 6, "xxxc", "c", ""},
		// Test "li", non-valid li letter preceeding
		{"xxxxli", 3, 6, "xxxxli", "xli", ""},
		// Test "ogi"
		{"xxlogi", 2, 6, "xxlog", "log", ""},
		// Test "ogi", not preceeded by "l"
		{"xxxogi", 2, 6, "xxxogi", "xogi", ""},
		// Test the others, which are simple replacements
		{"xxxxenci", 3, 7, "xxxxence", "xence", "e"},
		{"xxxxanci", 3, 7, "xxxxance", "xance", "e"},
		{"xxxxabli", 3, 7, "xxxxable", "xable", "e"},
		{"xxxxentli", 3, 8, "xxxxent", "xent", ""},
		{"xxxxizer", 3, 7, "xxxxize", "xize", ""},
		{"xxxxization", 3, 10, "xxxxize", "xize", ""},
		{"xxxxational", 3, 10, "xxxxate", "xate", ""},
		{"xxxxation", 3, 8, "xxxxate", "xate", ""},
		{"xxxxator", 3, 7, "xxxxate", "xate", ""},
		{"xxxxalism", 3, 8, "xxxxal", "xal", ""},
		{"xxxxaliti", 3, 8, "xxxxal", "xal", ""},
		{"xxxxalli", 3, 7, "xxxxal", "xal", ""},
		{"xxxxfulness", 3, 10, "xxxxful", "xful", ""},
		{"xxxxousli", 3, 8, "xxxxous", "xous", ""},
		{"xxxxousness", 3, 10, "xxxxous", "xous", ""},
		{"xxxxiveness", 3, 10, "xxxxive", "xive", ""},
		{"xxxxiviti", 3, 8, "xxxxive", "xive", ""},
		{"xxxxbiliti", 3, 9, "xxxxble", "xble", ""},
		{"xxxxbli", 3, 6, "xxxxble", "xble", "e"},
		{"xxxxfulli", 3, 8, "xxxxful", "xful", ""},
		{"xxxxlessli", 3, 8, "xxxxless", "xless", ""},
		// Some of the same words, this time not in our fake R1
		{"xxxxenci", 8, 8, "xxxxenci", "", ""},
		{"xxxxanci", 8, 8, "xxxxanci", "", ""},
		{"xxxxabli", 8, 8, "xxxxabli", "", ""},
		{"xxxxentli", 9, 9, "xxxxentli", "", ""},
		{"xxxxizer", 8, 8, "xxxxizer", "", ""},
		{"xxxxization", 11, 11, "xxxxization", "", ""},
		{"xxxxational", 11, 11, "xxxxational", "", ""},
		{"xxxxation", 9, 9, "xxxxation", "", ""},
		{"xxxxator", 8, 8, "xxxxator", "", ""},
	}

	runStepTest(t, testCases, step2)
}

func Test_step4(t *testing.T) {
	var testCases = []testCase{
		{"accumulate", 2, 5, "accumul", "cumul", "ul"},
		{"agreement", 2, 6, "agreement", "reement", "ent"},
	}

	runStepTest(t, testCases, step4)
}

func Test_step5(t *testing.T) {
	var testCases = []testCase{
		{"skate", 4, 5, "skate", "e", ""},
	}

	runStepTest(t, testCases, step5)
}

func Test_Stem(t *testing.T) {
	var testCases = []struct {
		in  string
		out string
	}{
		{"aberration", "aberr"},
		{"abruptness", "abrupt"},
		{"absolute", "absolut"},
		{"abated", "abat"},
		{"acclivity", "accliv"},
		{"accumulations", "accumul"},
		{"agreement", "agreement"},
		{"breed", "breed"},
		{"ape", "ape"},
		{"skating", "skate"},
		{"fluently", "fluentli"},
		{"ied", "ie"},
		{"ies", "ie"},
		{"because", "becaus"},
		{"above", "abov"},
	}

	stemmer := PorterStemmer{}
	stemmer.Init(false, false, make(map[string]bool))

	for _, tc := range testCases {
		stemmed, _ := stemmer.Stem([]rune(tc.in))
		if string(stemmed) != tc.out {
			t.Errorf("Expected %s to stem to %s, but got %s", tc.in, tc.out, string(stemmed))
		}
	}

}
