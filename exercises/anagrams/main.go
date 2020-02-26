package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagrams(word string, words []string) []string {
	var results []string
	for i := 0; i < len(words); i++ {
		if isAnagram(word, words[i]) {
			results = append(results, words[i])
		}
	}
	return results
}

func isAnagram(word1 string, word2 string) bool {
	return word1 != word2 && reverseStr(word1) == reverseStr(word2)
}

func reverseStr(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	items := []string{"enlists", "google", "inlets", "banana"}
	list := anagrams("listen", items)
	fmt.Println(list)
}
