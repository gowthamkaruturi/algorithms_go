package arrays

import (
	"fmt"
	"sort"
	"strings"
)

type word struct {
	word  string
	count int64
}

func WordSort(s string) []word {
	countMap := make(map[string]int64, 0)
	words := make([]word, 0)

	for _, s := range strings.Fields(s) {
		countMap[s] = countMap[s] + 1
	}
	for k, v := range countMap {
		words = append(words, word{
			word:  k,
			count: v,
		})
	}
	fmt.Println("words before any sort", words)
	sort.Slice(words[:], func(i, j int) bool {
		return len(words[i].word) < len(words[j].word) && words[i].word < words[j].word
	})
	fmt.Println("after length sort", words)

	// sort.Slice(words[:], func(i, j int) bool {
	// 	return words[i].word < words[j].word
	// })
	fmt.Println("after lexical sort", words)

	return words
}
