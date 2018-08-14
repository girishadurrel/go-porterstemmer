# Go Porter (2) Stemmer

This is a GO implementation of the porter stemmer 2 algorithm. I forked the porter stemmer algorithm from this [repo](https://github.com/magiczhao/go-porterstemmer) and then changed the code to reflect Porter 2. The Porter 2 implementation in this algorithm is directly based upon the code found in this [repo](https://github.com/kljensen/snowball). If you are looking for a Porter 2 implementation in GO please refer to the original [repo](https://github.com/kljensen/snowball). 

# USAGE

```
package main

import (
	"fmt"

	"github.com/girishadurrel/go-porterstemmer"
)

func main() {
	//word to be stemmed
	input := "dresses"
	//want to check the stemmers output at every step?
	//set this to true and printout the slice debugOutput
	collectDebugDetails := false
	//plural nouns *"usually"* ends with s, es, ves, ies.
	//if partialStem is set to true then only the nouns
	//should get stemmed. not 100%certain on this :(
	partialStem := false

	stemmer := porterstemmer2.PorterStemmer{}
	//the special case map is used to ignore the algorithm
	//and map stemmed output from a dictionary
	//example earring and ear both map to ear if the algorithm
	//is used. Thus, earring can be added to this map to
	//earring -> earring
	specialCaseMap := make(map[string]string)
	//init the stemmer
	stemmer.Init(collectDebugDetails, partialStem, specialCaseMap)

	//stemm the input, if you don't feel like  converting to runes,
	//you can always use StemString()
	output, debugOuput := stemmer.Stem([]rune(input))

	//print the stemmed output output
	fmt.Println(input + " > " + string(output))

	if collectDebugDetails {
		for _, item := range debugOuput {
			fmt.Println(item)
		}
	}
}
```
