package porterstemmer2

import (
	"bufio"
	"flag"
	"os"
	"strings"
	"testing"
)

var (
	specialFilePath = flag.String("filepath", "en_special_stemm.txt", "file with special stemm maps")
)

func Test_partialStemm(t *testing.T) {
	var testCases = []struct {
		in  string
		out string
	}{
		{"cars", "car"},
		{"dogs", "dog"},
		{"books", "book"},
		{"houses", "hous"},
		{"apples", "appl"},
		{"days", "day"},
		{"keys", "key"},
		{"guys", "guy"},
		{"donkeys", "donkey"},
		{"buses", "bus"},
		{"matches", "match"},
		{"dishes", "dish"},
		{"boxes", "box"},
		{"quizzes", "quiz"},
		{"leaves", "leaf"},
		{"wolves", "wolf"},
		{"lives", "life"},
		{"knives", "knife"},
		{"roofs", "roof"},
		{"cliffs", "cliff"},
		{"boys", "boy"},
		{"cities", "citi"},
		{"babies", "babi"},
		{"stories", "stori"},
		{"parties", "parti"},
		{"countries", "countri"},
		{"zoos", "zoo"},
		{"radios", "radio"},
		{"stereos", "stereo"},
		{"videos", "video"},
		{"kangaroos", "kangaroo"},
		{"heroes", "hero"},
		{"echoes", "echo"},
		{"tomatoes", "tomato"},
		{"pianos", "piano"},
		{"potatoes", "potato"},
		{"series", "seri"},
		{"species", "speci"},
		{"berries", "berri"},
		//should stemm to "activ" on partial stemm, but need step3 to be active.
		//therefore left as a failed test as activing step3 would cause non nouns
		//to be stemmed. therefore will keep it as it is. the output will
		//therefore be activiti
		// -------------------- failed test case -------------------- //
		{"activities", "activ"},
		// -------------------- failed test case -------------------- //
		{"daisies", "daisi"},
		{"churches", "church"},
		{"foxes", "fox"},
		{"stomachs", "stomach"},
		{"epochs", "epoch"},
		{"halves", "half"},
		{"scarves", "scarf"},
		{"chiefs", "chief"},
		{"spoofs", "spoof"},
		{"solos", "solo"},
		{"zeros", "zero"},
		{"avocados", "avocado"},
		{"studios", "studio"},
		{"embryos", "embryo"},
		{"buffaloes", "buffalo"},
		{"dominoes", "domino"},
		{"embargoes", "embargo"},
		{"mosquitoes", "mosquito"},
		{"torpedoes", "torpedo"},
		{"vetoes", "veto"},
		{"espressos", "espresso"},
	}

	specialMap := make(map[string]string)

	dataFile, err := os.Open(*specialFilePath)

	if err != nil {
		panic(err)
	}

	defer dataFile.Close()

	reader := bufio.NewReader(dataFile)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		tks := strings.Split(text, "\t")

		if _, ok := specialMap[tks[0]]; !ok {
			specialMap[tks[0]] = tks[1]
		}
	}

	fullStemmer := PorterStemmer{}
	fullStemmer.Init(true, false, specialMap)

	partialStemmer := PorterStemmer{}
	partialStemmer.Init(false, true, specialMap)

	for _, tc := range testCases {
		fstemmed, _ := fullStemmer.Stem([]rune(tc.in))
		pstemmed, _ := partialStemmer.Stem([]rune(tc.in))

		if string(fstemmed) != string(pstemmed) {
			t.Errorf("given %s: full stemm gives > %s while partail stemm gives > %s", tc.in, string(fstemmed), string(pstemmed))
		}
	}

}
